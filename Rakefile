task :default => [:build]

task build: [ :go_build ] do
end

task go_generate: [] do
    sh "goagen bootstrap -o src/ -d github.com/alligrader/gradebook-api/designs"
end

task go_build: [ :go_generate ] do
    sh "go build -o main ./src"
end

task clean: [] do
    sh "rm main"
end
