package jMiddleware

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/modules/j/jConstants"
	"github.com/ZhanibekTau/go-jm-core/pkg/modules/j/jStructures"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func RequestHandler(requestData *jStructures.RequestData) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.GetHeader(jConstants.UserIdHeaderName))
		if err != nil {
			userId = 0
		}

		companyId, err := strconv.Atoi(c.GetHeader(jConstants.CompanyIdHeaderName))
		if err != nil {
			companyId = 0
		}

		currentCompanyId, err := strconv.Atoi(c.GetHeader(jConstants.CurrentCompanyIdHeaderName))
		if err != nil {
			currentCompanyId = 0
		}

		companyIdsStr := c.GetHeader(jConstants.CompanyIdsHeaderName)
		arr := strings.Split(companyIdsStr, ",")

		companyIds := make([]int, len(arr))

		for i, v := range arr {
			companyIds[i], _ = strconv.Atoi(v)
		}

		requestData.CompanyIds = companyIds
		requestData.UserId = userId
		requestData.UserType = c.GetHeader(jConstants.UserTypeHeaderName)
		requestData.RequestId = c.GetHeader(jConstants.RequestIdHeaderName)
		requestData.LanguageCode = c.GetHeader(jConstants.LanguageCodeHeaderName)
		requestData.CompanyId = companyId
		requestData.CurrentCompanyId = currentCompanyId
		requestData.RequestUrl = c.Request.URL.Path
		requestData.RequestMethod = c.Request.Method
		requestData.RequestScheme = c.Request.URL.Scheme
		requestData.RequestHost = c.Request.Host

		c.Next()
	}
}
