package helpers

import "os"

func BuildForPython() error {
	err := Render()
	if err != nil {
		return err
	}

	err = os.CopyFS("build/python", os.DirFS("codes/python"))
	if err != nil {
		return err
	}

	err = os.CopyFS("build/python/build", os.DirFS("build/dev"))
	if err != nil {
		return err
	}

	return nil
}

func BuildForNode() error {
	err := Render()
	if err != nil {
		return err
	}

	err = os.CopyFS("build/node", os.DirFS("codes/node"))
	if err != nil {
		return err
	}

	err = os.CopyFS("build/node/src/build", os.DirFS("build/dev"))
	if err != nil {
		return err
	}

	return nil
}

func BuildForHtml() error {
	err := Render()
	if err != nil {
		return err
	}

	err = os.CopyFS("build/html", os.DirFS("build/dev"))
	if err != nil {
		return err
	}

	return nil
}
