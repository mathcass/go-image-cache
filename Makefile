# Simple Makefile to aid common operations

serve:
	goapp serve

deploy:
	goapp deploy

clean:
	goapp serve -clear_datastore
