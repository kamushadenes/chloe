module github.com/kamushadenes/chloe

go 1.21

replace github.com/ggerganov/whisper.cpp/bindings/go => ./workspace/models/audio_models/whisper.cpp/bindings/go

require (
	cloud.google.com/go/texttospeech v1.7.1
	github.com/MichaelMure/go-term-markdown v0.1.4
	github.com/alecthomas/kong v0.8.0
	github.com/antonmedv/expr v1.15.2
	github.com/aquilax/truncate v1.0.0
	github.com/biter777/countries v1.6.6
	github.com/briandowns/spinner v1.23.0
	github.com/bwmarrin/discordgo v0.27.1
	github.com/fatih/color v1.15.0
	github.com/ggerganov/whisper.cpp/bindings/go v0.0.0-20230912105404-3fec2119e6b5
	github.com/go-audio/wav v1.1.0
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/render v1.0.3
	github.com/go-latex/latex v0.0.0-20230307184459-12ec69307ad9
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/gocolly/colly v1.2.0
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/haguro/elevenlabs-go v0.2.2
	github.com/iancoleman/strcase v0.3.0
	github.com/jedib0t/go-pretty/v6 v6.4.7
	github.com/mattn/go-isatty v0.0.19
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5
	github.com/pkg/profile v1.7.0
	github.com/pkoukk/tiktoken-go v0.1.6
	github.com/rocketlaunchr/google-search v1.1.6
	github.com/rs/zerolog v1.30.0
	github.com/sashabaranov/go-openai v1.15.3
	github.com/slack-go/slack v0.12.3
	github.com/stretchr/testify v1.8.4
	github.com/trietmn/go-wiki v1.0.1
	golang.org/x/term v0.12.0
	google.golang.org/api v0.143.0
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.5.3
	gorm.io/driver/sqlserver v1.5.1
	gorm.io/gorm v1.25.4
	mvdan.cc/xurls/v2 v2.5.0
)

require (
	cloud.google.com/go v0.110.8 // indirect
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/longrunning v0.5.1 // indirect
	github.com/MichaelMure/go-term-text v0.3.1 // indirect
	github.com/PuerkitoBio/goquery v1.8.1 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/alecthomas/chroma v0.10.0 // indirect
	github.com/anaskhan96/soup v1.2.5 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/antchfx/htmlquery v1.3.0 // indirect
	github.com/antchfx/xmlquery v1.3.17 // indirect
	github.com/antchfx/xpath v1.2.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/eliukblau/pixterm/pkg/ansimage v0.0.0-20191210081756-9fb6cf8c2f75 // indirect
	github.com/felixge/fgprof v0.9.3 // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/go-audio/audio v1.0.0 // indirect
	github.com/go-audio/riff v1.0.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly/v2 v2.1.0 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gomarkdown/markdown v0.0.0-20230912175223-14b07df9d538 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20230912144702-c363fe2c2ed8 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.1 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kyokomi/emoji/v2 v2.2.12 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/microsoft/go-mssqldb v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/saintfish/chardet v0.0.0-20230101081208-5e3ef4b5456d // indirect
	github.com/temoto/robotstxt v1.1.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/image v0.12.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/oauth2 v0.12.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20230913181813-007df8e322eb // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230913181813-007df8e322eb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/grpc v1.58.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
