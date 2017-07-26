require 'rake'

task :test do
    sh("go test ./... -v")
end

task :cover do
    sh("go test --coverprofile=coverage.out")
    sh("go tool cover --html=coverage.out")
end
