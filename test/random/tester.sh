for i in {0..100}
do
    go run randomizer.go 9
    cat "test.txt"
    time go run ../../app/main.go "test.txt"
done