cur_path=$(pwd)
export GOPATH=${cur_path}

go build ./ -o ./deployment/script