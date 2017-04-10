testrun:
	go build -o test_workflow
	echo `cat test/jobs.yaml`
	./test_workflow run "`cat test/jobs.yaml`"
	rm test_workflow
