setup:
	pip3 install virtualenv
	pip3 install wheel
	pip3 install setuptools
	pip3 install twine

install:
	pip3 install twine
	pip3 install setuptools
	pip3 install wheel
	pip3 install -r requirements.txt
	

release:
	python3 setup.py sdist bdist_wheel
	twine upload --repository testpypi --skip-existing dist/*

install-release:
	pip install --index-url https://test.pypi.org/simple/ mob-devely