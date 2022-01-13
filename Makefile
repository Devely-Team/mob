setup:
	pip3 install virtualenv
	pip3 install wheel
	pip3 install setuptools
	pip3 install twine
	pip3 install -r requirements.txt

install:
	virtualenv devely
	source devely/bin/activate
	python3 setup.py develop

release:
	python3 setup.py sdist bdist_wheel
	twine upload --repository testpypi --skip-existing dist/*

install-release:
	pip install --index-url https://test.pypi.org/simple/ mob-devely

clean:
	rm -rf dist
	rm -rf build
	rm -rf devely
	rm -rf mob-devely.egg-info