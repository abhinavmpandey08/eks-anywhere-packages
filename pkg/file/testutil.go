package file

import (
	"github.com/aws/eks-anywhere-packages/api/v1alpha1"
)

func givenFile(file string, config KindAccessor) error {
	reader := NewFileReader(file)
	err := reader.Initialize(config)
	if err != nil {
		return err
	}
	return reader.Parse(config)
}

func GivenPackage(fileName string) (*v1alpha1.Package, error) {
	config := &v1alpha1.Package{}
	err := givenFile(fileName, config)
	return config, err
}

func GivenPackageBundle(filename string) (*v1alpha1.PackageBundle, error) {
	config := &v1alpha1.PackageBundle{}
	reader := NewFileReader(filename + ".signed")
	initError := reader.Initialize(config)
	if initError != nil {
		return nil, initError
	}
	actual := reader.Parse(config)
	if actual != nil {
		return nil, actual
	}
	return config, nil
}

func GivenBundleController(fileName string) (*v1alpha1.PackageBundleController, error) {
	config := &v1alpha1.PackageBundleController{}
	err := givenFile(fileName, config)
	return config, err
}
