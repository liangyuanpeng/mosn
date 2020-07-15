function make_checker {
	go build -o tokener tokener.go
}

function make_mosn {
	mkdir ./build_mosn
	cp ../../../../../cmd/mosn/main/* ./build_mosn
	cp ./tokener_filter.go ./build_mosn
	cd ./build_mosn
	go build -o mosn
	mv mosn ../
	cd ../
	rm -rf ./build_mosn
}

make_checker
make_mosn
