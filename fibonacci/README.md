$ go test -bench=. -count=10 | tee old.txt

$ go test -bench=. -count=10 | tee new.txt
