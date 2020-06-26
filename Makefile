NAME=terraform-provider-circleci
PLUGIN_PATH=$(HOME)/.terraform.d/plugins

all: build

build:
	go build -o $(NAME)


install: build
		install -d $(PLUGIN_PATH)
		install -m 775 $(NAME) $(PLUGIN_PATH)/
