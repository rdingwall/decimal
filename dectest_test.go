package decimal_test

import (
	"path/filepath"
	"testing"

	"github.com/ericlagergren/decimal/dectest"
)

func TestDecTests(t *testing.T) {
	files, err := filepath.Glob("_dectest/*.decTest")
	if err != nil {
		t.Fatal(err)
	}

	if len(files) == 0 {
		t.Skip("No .detect files found, please run _dectests/generate.bash")
	}

	//files = []string{"_dectest/log10.decTest"}
	/* files = []string{"_dectest/log10.decTest"}
	files = []string{
		"_dectest/abs.decTest",
		"_dectest/add.decTest",
		"_dectest/clamp.decTest",
		"_dectest/class.decTest",
		"_dectest/compare.decTest",
		"_dectest/copy.decTest",
		"_dectest/copysign.decTest",
		"_dectest/ddAbs.decTest",
		"_dectest/ddCanonical.decTest",
		"_dectest/ddClass.decTest",
		"_dectest/ddCompare.decTest",
		"_dectest/ddCopy.decTest",
		"_dectest/ddCopySign.decTest",
		"_dectest/ddDivide.decTest",
		"_dectest/ddDivideInt.decTest",
		"_dectest/ddFMA.decTest",
		"_dectest/ddLogB.decTest",
		"_dectest/ddMax.decTest",
		"_dectest/ddMin.decTest",
		"_dectest/ddMinus.decTest",
		"_dectest/ddMultiply.decTest",
		"_dectest/ddQuantize.decTest",
		"_dectest/ddReduce.decTest",
		"_dectest/ddRemainder.decTest",
		"_dectest/ddShift.decTest",
		"_dectest/ddSubtract.decTest",
		"_dectest/ddToIntegral.decTest",
		"_dectest/divide.decTest",
		"_dectest/divideInt.decTest",
		"_dectest/fma.decTest",
		"_dectest/ln.decTest",
		"_dectest/quantize.decTest",
		"_dectest/rounding.decTest",
		"_dectest/squareroot.decTest",
		"_dectest/subtract.decTest",
	} */

	for _, file := range files {

		/*if strings.Contains(file, "ncode") || strings.Contains(file, "anonical") {
			continue
		}*/
		file := file // shadow range variable
		t.Run(filepath.Base(file), func(t *testing.T) {
			dectest.Test(t, file)
		})
	}
}
