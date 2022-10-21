package di

import (
	"log"

	"github.com/jphacks/B_2206/api/drivers/db"
	"github.com/jphacks/B_2206/api/drivers/server"
	"github.com/jphacks/B_2206/api/externals/controller"
	"github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/jphacks/B_2206/api/router"
)

func InitializeServer() db.Client {
	// DB接続
	client, err := db.ConnectMySQL()
	if err != nil {
		log.Fatal("db error")
	}

	// ↓

	// Repository
	userRepository := repository.NewUserRepository(client)
	mailAuthRepository := repository.NewMailAuthRepository(client)
	areaRepository := repository.NewAreaRepository(client)
	classificationRepository := repository.NewClassificationRepository(client)
	companyInfoRepository := repository.NewCompanyInfoRepository(client)
	detailRepository := repository.NewDetailRepository(client)
	limitRepository := repository.NewLimitRepository(client)
	matchingRepository := repository.NewMatchingRepository(client)
	personalInfoRepository := repository.NewPersonalInfoRepository(client)
	requestRepository := repository.NewRequestRepository(client)
	tagRepository := repository.NewTagRepository(client)
	valueRepository := repository.NewValueRepository(client)
	watchListRepository := repository.NewWatchListRepository(client)
	wishListRepository := repository.NewWishListRepository(client)
	// ↓

	// UseCase
	userUseCase := usecase.NewUserUseCase(userRepository, areaRepository)
	mailAuthUseCase := usecase.NewMailAuthUseCase(mailAuthRepository, areaRepository)
	classificationUseCase := usecase.NewClassificationUseCase(classificationRepository)
	companyInfoUseCase := usecase.NewCompanyInfoUseCase(companyInfoRepository)
	detailUseCase := usecase.NewDetailUseCase(detailRepository)
	limitUseCase := usecase.NewLimitUseCase(limitRepository)
	matchingUseCase := usecase.NewMatchingUseCase(matchingRepository)
	personalInfoUseCase := usecase.NewPersonalInfoUseCase(personalInfoRepository)
	requestUseCase := usecase.NewRequestUseCase(requestRepository)
	tagUseCase := usecase.NewTagUseCase(tagRepository)
	valueUseCase := usecase.NewValueUseCase(valueRepository)
	watchListUseCase := usecase.NewWatchListUseCase(watchListRepository)
	wishListUseCase := usecase.NewWishListUseCase(wishListRepository)

	// ↓

	// Controller
	healthcheckController := controller.NewHealthCheckController()
	mailAuthController := controller.NewMailAuthController(mailAuthUseCase)
	userController := controller.NewUserController(userUseCase)
	classificationController := controller.NewClassificationController(classificationUseCase)
	companyInfoController := controller.NewCompanyInfoController(companyInfoUseCase)
	detailController := controller.NewDetailController(detailUseCase)
	limitController := controller.NewLimitController(limitUseCase)
	matchingController := controller.NewMatchingController(matchingUseCase)
	personalInfoController := controller.NewPersonalInfoController(personalInfoUseCase)
	requestController := controller.NewRequestController(requestUseCase)
	tagController := controller.NewTagController(tagUseCase)
	valueController := controller.NewValueController(valueUseCase)
	watchListController := controller.NewWatchListController(watchListUseCase)
	wishListController := controller.NewWishListController(wishListUseCase)

	// ↓

	// router
	router := router.NewRouter(
		healthcheckController,
		mailAuthController,
		userController,
		classificationController,
		companyInfoController,
		detailController,
		limitController,
		matchingController,
		personalInfoController,
		requestController,
		tagController,
		valueController,
		watchListController,
		wishListController,
	)

	// ↓

	// Server
	server.RunServer(router)

	return client
}
