task :default => [:all]

task build: [ :go_build ] do
end

task all: [ :go_generate, :build ] do

end

task go_generate: [] do
    sh "goagen bootstrap -o src/ -d github.com/alligrader/gradebook-api/designs"
    sh "goagen gen -o src/       -d github.com/alligrader/gradebook-api/designs --pkg-path=github.com/goadesign/gorma"
    sh "goagen gen               -d github.com/alligrader/gradebook-api/designs --pkg-path=github.com/kkeuning/reduxa"
    sh "swagger2blueprint src/swagger/swagger.json apiary.apib"
end

task go_build: [] do
    sh "go build -o main ./src"
end

task clean: [] do
    sh "rm main"
end

task k8s: [] do
    sh "kubectl proxy --address='0.0.0.0'"
end
