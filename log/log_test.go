package log

import (
	"testing"

	"github.com/go-angle/angle"
	"github.com/go-angle/angle/config"
)

func BenchmarkLogConsoleContextFields(b *testing.B) {
	angle.RunOnce("config.test.yml", func(l Logger) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l.WithFields(Fields{
				"int":    1,
				"string": "string",
				"f":      1.0,
				"map": Fields{
					"123": 123,
					"s":   "s",
				},
				"arr": []Fields{
					Fields{
						"s": "s",
						"i": 10,
					},
				},
			}).Infof("%s", "test")
		}
	})
}

func BenchmarkLogJSONContextFields(b *testing.B) {
	angle.RunOnce("config.prod.yml", func(c *config.Config, l Logger) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			l.WithFields(Fields{
				"int":    1,
				"string": "string",
				"f":      1.0,
				"map": Fields{
					"123": 123,
					"s":   "s",
				},
				"arr": []Fields{
					Fields{
						"s": "s",
						"i": 10,
					},
				},
			}).Infof("%s", "test")
		}
	})

}

func TestLog(t *testing.T) {
	angle.RunOnce("config.test.yml", func(log Logger) {
		log.Info("without extra fields")
		log.WithFields(Fields{
			"fields1": 1,
		}).Info("without extra fields1")
		l := log.WithFields(Fields{
			"fields2": 1,
		})
		l.Info("without extra fields2")
		l.Info("without extra fields2")

		l = log.WithFields(Fields{
			"fields2": 2,
		})
		l.Info("without extra fields2")
	})
}
