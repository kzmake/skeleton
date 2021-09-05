.PHONY: all
all: pre fmt lint

.PHONY: pre
pre:
	poetry install --no-root

.PHONY: fmt
fmt:
	black .
	isort .

.PHONY: lint
lint:
	flake8

.PHONY: run
run:
	PYTHONPATH=../../../gen/py/ python run.py
