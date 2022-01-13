from setuptools import setup, find_packages
with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()
with open("requirements.txt", "r", encoding="utf-8") as fh:
    requirements = fh.read()
setup(
    name='mob-devely',
    version='0.0.4',
    author='Devely',
    author_email='develyteam@gmail.com',
    license='MIT',
    description='A flutter project manager.',
    long_description_content_type="text/markdown",
    url='https://github.com/Devely-Team/mob',
    py_modules=['mob', 'cli'],
    packages=find_packages(),
    install_requires=[requirements],
    python_requires='>=3.7',
    classifiers=[
        "Programming Language :: Python :: 3.9",
        "Operating System :: OS Independent",
    ],
    entry_points='''
        [console_scripts]
        mob=mob:cli
    '''
)
