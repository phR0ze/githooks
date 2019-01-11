NAME := githooks
DEPDIR := vendor
PKGROOT := github.com/phR0ze/githooks

default: $(NAME)
$(NAME): $(DEPDIR)
	go build -o build/$(NAME) $(PKGROOT)/cmd/githooks

install:
	dep ensure -v

update:
	dep ensure -v -update

test: $(NAME)
	@echo -e "\nRunning all go tests:"
	@echo -e "------------------------------------------------------------------------"
	go test $(PKGROOT)/cmd/githooks

bench: $(NAME)
	@echo -e "\nRunning all go benchmarks:"
	@echo -e "------------------------------------------------------------------------"
	go test -bench=.

clean:
	rm -rf ./vendor
	rm -f ./bin/$(NAME)
