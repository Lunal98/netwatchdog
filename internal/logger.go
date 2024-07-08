/*
Copyright © 2024 Alex Bedo <alex98hun@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Setup() {
	multi := zerolog.MultiLevelWriter(os.Stdout, rotate("pwdgen.log"))
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

}
func rotate(path string) io.WriteCloser {
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    200,
		MaxAge:     90,
		MaxBackups: 10,
		Compress:   true,
	}
}
