package server

import (
	"ace/controller"
	"ace/middleware"
	"github.com/gin-gonic/gin"
)

func NewServer(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.New()

	r.Use(middleware.Cors(), middleware.Logger(), middleware.Recovery())

	api := r.Group("/api/v1")
	{
		api.POST("/login", controller.LoginController)
		api.POST("/register", controller.RegisterController)

		user := api.Group("/user")
		user.Use(middleware.Auth())
		{
			user.GET("/getUserInfo", controller.GetUserInfoController)
		}

		common := api.Group("/common")
		common.Use(middleware.Auth())
		{
			common.GET("/platforms", controller.CommonPlatforms)
			common.GET("/regions", controller.CommonRegion)
			common.GET("/tones", controller.CommonTones)
			common.GET("/voices", controller.CommonVoices)
			common.POST("/createVoice", controller.CommonCreateVoice)
			common.GET("/languages", controller.CommonLanguages)
			common.GET("/genders", controller.CommonGender)
		}

		content := api.Group("/content")
		content.Use(middleware.Auth())
		{
			media := content.Group("/media")
			{
				media.POST("/generator", controller.MediaGenerator)
				media.GET("/histories", controller.MediaHistoryController)
				media.GET("/getHistoryById", controller.MediaGetHistoryById)
			}

			engine := content.Group("/engine")
			{
				engine.POST("/generator", controller.EngineGenerator)
				engine.GET("/histories", controller.EngineHistoryController)
				engine.GET("/getHistoryById", controller.EngineGetHistoryById)
			}

			optimized := content.Group("/optimized")
			{
				tone := optimized.Group("/tone")
				{
					tone.POST("/generator", controller.ToneGenerator)
					tone.GET("/histories", controller.ToneHistory)
					tone.GET("/getHistoryById", controller.ToneHistoryById)
				}

				summarize := optimized.Group("/summarize")
				{
					summarize.POST("/generator", controller.SummarizeGenerator)
					summarize.GET("/histories", controller.SummarizeHistory)
					summarize.GET("/getHistoryById", controller.SummarizeHistoryById)
				}

				paraphrase := optimized.Group("/paraphrase")
				{
					paraphrase.POST("/generator", controller.ParaphraseGenerator)
					paraphrase.GET("/histories", controller.ParaphraseHistory)
					paraphrase.GET("/getHistoryById", controller.ParaphraseHistoryById)
				}

				voice := optimized.Group("/voice")
				{
					voice.POST("/generator", controller.VoiceGenerator)
					voice.GET("/histories", controller.VoiceHistory)
					voice.GET("/getHistoryById", controller.VoiceHistoryById)
				}

				audience := optimized.Group("/audience")
				{
					audience.POST("/generator", controller.AudienceGenerator)
					audience.GET("/histories", controller.AudienceHistory)
					audience.GET("/getHistoryById", controller.AudienceHistoryById)
				}
			}

			email := content.Group("/email")
			{
				freestyle := email.Group("/freestyle")
				{
					freestyle.POST("/generator", controller.FreestyleGenerator)
					freestyle.GET("/histories", controller.FreestyleHistory)
					freestyle.GET("/getHistoryById", controller.FreestyleHistoryById)
				}

				marketing := email.Group("/marketing")
				{
					marketing.POST("/generator", controller.MarketingGenerator)
					marketing.GET("/histories", controller.MarketingHistory)
					marketing.GET("/getHistoryById", controller.MarketingHistoryById)
				}

				welcome := email.Group("/welcome")
				{
					welcome.POST("/generator", controller.WelcomeGenerator)
					welcome.GET("/histories", controller.WelcomeHistory)
					welcome.GET("/getHistoryById", controller.WelcomeHistoryById)
				}

				advantage := email.Group("/advantage")
				{
					advantage.POST("/generator", controller.AdvantageGenerator)
					advantage.GET("/histories", controller.AdvantageHistory)
					advantage.GET("/getHistoryById", controller.AdvantageHistoryById)
				}
			}

			blog := content.Group("/blog")
			{
				intro := blog.Group("/intro")
				{
					intro.POST("/generator", controller.IntroGenerator)
					intro.GET("/histories", controller.IntroHistory)
					intro.GET("/getHistoryById", controller.IntroHistoryById)
				}

				outline := blog.Group("/outline")
				{
					outline.POST("/generator", controller.OutlineGenerator)
					outline.GET("/histories", controller.OutlineHistory)
					outline.GET("/getHistoryById", controller.OutlineHistoryById)
				}

				entire := blog.Group("/entire")
				{
					entire.POST("/generator", controller.EntireGenerator)
					entire.GET("/histories", controller.EntireHistory)
					entire.GET("/getHistoryById", controller.EntireHistoryById)
				}
			}
		}
	}

	return r
}
