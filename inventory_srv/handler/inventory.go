package handler

import (
	"context"
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	//"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	//"gorm.io/gorm/clause"
	"hcshop_srvs/inventory_srv/global"
	"hcshop_srvs/inventory_srv/model"
	"hcshop_srvs/inventory_srv/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (i InventoryServer) SetInv(ctx context.Context, in *proto.GoodsInvInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: in.GoodsId}).First(&inv)
	inv.Goods = in.GoodsId
	inv.Stocks = in.Num

	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}

func (i InventoryServer) InvDetail(ctx context.Context, in *proto.GoodsInvInfo, opts ...grpc.CallOption) (*proto.GoodsInvInfo, error) {
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: in.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到该商品的库存信息")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil

}

func (i InventoryServer) Sell(ctx context.Context, in *proto.SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	// 这里使用分布式锁来实现
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
		Password: global.ServerConfig.RedisInfo.Password,
	})
	pools := goredis.NewPool(client) // 这里可以使用多个client 初始化
	rs := redsync.New(pools)

	tx := global.DB.Begin()
	sellDetail := model.StockSellDetail{
		OrderSn: in.OrderSn,
		Status:  1,
	}
	var goodsInvInfos []model.GoodsDetail // 所有要扣的商品ID和对应的要扣件的库存
	for _, goodsInvInfo := range in.GoodsInfo {
		goodsInvInfos = append(goodsInvInfos, model.GoodsDetail{
			Goods: goodsInvInfo.GoodsId,
			Num:   goodsInvInfo.Num,
		})
		// 这是要扣减的库存对象
		var inv model.Inventory
		//搞个分布式同步锁出来，使用商品ID作为key
		mutext := rs.NewMutex(fmt.Sprintf("goods_%d", goodsInvInfo.GoodsId))
		// 获取锁
		if err := mutext.Lock(); err != nil {
			return nil, status.Errorf(codes.Internal, "获取分布式锁失败")
		}

		if result := global.DB.Where(&model.Inventory{Goods: goodsInvInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "没有该商品库存信息")
		}

		if inv.Stocks < goodsInvInfo.Num {
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}
		inv.Stocks -= goodsInvInfo.Num
		tx.Save(&inv)

		if ok, err := mutext.Unlock(); !ok || err != nil {
			return nil, status.Errorf(codes.Internal, "释放redis 分布式锁失败")
		}
	}
	sellDetail.Detail = goodsInvInfos
	if result := tx.Create(&sellDetail); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "保存库存扣减历史失败")
	}
	tx.Commit()
	return &emptypb.Empty{}, nil

	//tx := global.DB.Begin() // 这是数据级别的事物
	//sellDetail := model.StockSellDetail{
	//	OrderSn: in.OrderSn,
	//	Status:  1,
	//}
	//
	////
	//var goodsInvInfos []model.GoodsDetail // 所有要扣的商品ID和对应的要扣件的库存
	//for _, goodsInvInfo := range in.GoodsInfo {
	//	goodsInvInfos = append(goodsInvInfos, model.GoodsDetail{
	//		Goods: goodsInvInfo.GoodsId,
	//		Num:   goodsInvInfo.Num,
	//	})
	//
	//	var inv model.Inventory
	//
	//	// 这个是添加行锁，也就是mysql的 `  ....  for update`
	//	// 其实就是悲观锁，每次准备做修改，就去获取锁，只有获取了才做操作，保证我在修改库存的是时候，其他人不能修改我这个商品的库存
	//	//if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goodsInvInfo.GoodsId}).First(&inv);result.RowsAffected == 0{
	//	//	tx.Rollback()
	//	//	return nil,status.Errorf(codes.InvalidArgument,"没有库存信息")
	//	//}
	//
	//	// 这里是乐观锁的实现，通过失败重试来完成最终的目的，如果我成功了，我就不需要处理了。失败了，那就我继续再试，知道成功
	//	// 这里个人觉得应该有问题，难道要一直重试？ 是不是应该像redis或者mysql一样设置最长重试时间或者连接次数？
	//	for {
	//		if result := global.DB.Where(&model.Inventory{Goods: goodsInvInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
	//			tx.Rollback()
	//			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
	//		}
	//
	//		if inv.Stocks < goodsInvInfo.Num {
	//			tx.Rollback()
	//			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
	//		}
	//		// 因为这个操作不是原子操作，所以才要保护起来，但这个保护也只是在单机上，如果是集群环境，则会出现扣减不一致问题
	//		inv.Stocks = inv.Stocks - goodsInvInfo.Num
	//		//tx.Save(&inv)
	//
	//		if result := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods =? and version = ?", inv.Goods, inv.Version).Updates(model.Inventory{
	//			Stocks:  inv.Stocks,
	//			Version: inv.Version + 1,
	//		}); result.RowsAffected == 0 {
	//			zap.S().Info("库存扣减失败")
	//		} else {
	//			// 成功了就不需要重复了
	//			break
	//		}
	//	}
	//	tx.Save(&inv)
	//}
	//sellDetail.Detail = goodsInvInfos
	//if result := tx.Create(&sellDetail); result.RowsAffected == 0 {
	//	tx.Rollback()
	//	return nil, status.Errorf(codes.Internal, "保存库存扣减历史失败")
	//}
	//tx.Commit()
	//return &emptypb.Empty{}, nil
}

func (i InventoryServer) Reback(ctx context.Context, in *proto.SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}
