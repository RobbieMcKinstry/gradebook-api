task :default => [:build]

task build: [ :go_build ] do
end

task go_generate: [] do
    sh "goagen bootstrap -o src/ -d github.com/alligrader/gradebook-api/designs"
    sh "goagen gen               -d github.com/alligrader/gradebook-api/designs --pkg-path=github.com/goadesign/gorma"
    sh "goagen gen               -d github.com/alligrader/gradebook-api/designs --pkg-path=github.com/kkeuning/reduxa"
end

task go_build: [ :go_generate ] do
    sh "go build -o main ./src"
end

task clean: [] do
    sh "rm main"
end
