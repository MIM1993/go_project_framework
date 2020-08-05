package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go_project_framework/utils"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//双写,可以将返回的数据保存一份,用于打印log
func (w AccessLogWriter) Writer(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func LogMiddleware(logger *logrus.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		bodyWriter := AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}

		c.Writer = bodyWriter

		reqID := c.Request.Header.Get("X-Request-Id")
		if reqID == "" {
			reqID = c.Request.Header.Get("request_id")
		}
		if reqID == "" {
			reqID = utils.GenLogId()
		}
		realIP := c.Request.Header.Get("X-Real-IP")

		logEntry := logger.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"url_path":   c.Request.URL.Path,
			"client":     c.Request.RemoteAddr,
			"request_id": reqID,
			"real_ip":    realIP,
		})

		c.Set("logEntry", logEntry)
		c.Set("request_id", reqID)
		//var err error
		//c.Set("error", err)

		begin := time.Now()
		c.Set("st", begin)

		c.Next()
		//if cerr, ok := c.Get("error"); ok && cerr != nil {
		//	if ierr, ok := cerr.(error); ok || ierr == nil {
		//		err = ierr
		//	} else {
		//		err = fmt.Errorf("err: %v", cerr)
		//	}
		//}

		//c.MustGet("logEntry").(*logrus.Entry).WithFields(logrus.Fields{
		//	"total_cost": time.Now().Sub(begin).Seconds(),
		//}).Infof("finish. err:%v", err)
	}
}
