from setuptools import setup, find_packages

__version__ = "1.0.0"

setup(
    name='nebulaidl',
    version=__version__,
    description='IDL for Nebula Platform',
    url='https://www.github.com/nebulaclouds/nebulaidl',
    maintainer='NebulaClouds',
    maintainer_email='admin@nebula.org',
    packages=find_packages('gen/pb_python'),
    package_dir={'': 'gen/pb_python'},
    package_data={'nebulaidl': ["*.pyi", "**/*.pyi"]},
    dependency_links=[],
    install_requires=[
        'googleapis-common-protos',
        'protoc_gen_swagger',
        'protobuf>=4.21.1,<5.0.0',
    ],
    extras_require={
        ':python_version=="2.7"': ['typing>=3.6'],  # allow typehinting PY2
    },
)
