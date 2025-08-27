package main

import "github.com/joseluis244/file2cloud"

func main() {
	file2cloud.Init("391335f0facc50126a0c91545c69a574", "dbc92a6d8b73782ad99db62b013d6c7bd7288c4f206f0db6985a60107da4350c", "https://d819815835fcab8b93773d68a24d2722.r2.cloudflarestorage.com", "auto", "oncoservice")
	file2cloud.Upload("/Users/josecamacho/Documents/medicaresoft/cloudtest/DCM/8b/14/8b142869-56f4-4495-ac09-04ed3d8ecd06", "of1/study/serie/test.dcm", "application/dicom")
}
