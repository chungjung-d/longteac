package process

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/chungjung-d/longteac/config"
)

func StartContainerRootProcess(ctx context.Context) {
	fmt.Printf("New namespace setup complete, running as PID %d\n", os.Getpid())

	containerDirPath := ctx.Value(config.ContainerDirPath).(string)

	err := os.Chdir(containerDirPath)
	if err != nil {
		fmt.Println("Error changing directory: ")
		log.Fatal(err)
	}

	ociConfig, err := config.GetOCIConfig(containerDirPath)
	if err != nil {
		fmt.Println("Error getting OCI config: ")
		log.Fatal(err)
	}

	err = settingOverlayFsMount(containerDirPath)
	if err != nil {
		fmt.Println("Error setting overlay fs mount: ")
		log.Fatal(err)
	}

	err = settingMount(containerDirPath, ociConfig)
	if err != nil {
		fmt.Println("Error setting mount: ")
		log.Fatal(err)
	}

	err = settingPivotRoot(containerDirPath)
	if err != nil {
		fmt.Println("Error setting pivot root: ")
		log.Fatal(err)
	}

	// test_cmd := exec.Command("/bin/bash")
	// test_cmd.Stdout = os.Stdout
	// test_cmd.Stderr = os.Stderr
	// test_cmd.Stdin = os.Stdin
	// if err := test_cmd.Run(); err != nil {
	// 	fmt.Printf("Error listing /bin directory - %s\n", err)
	// 	os.Exit(1)
	// }

	cmd := exec.Command(ociConfig.ProcessConfig.Args[0], ociConfig.ProcessConfig.Args[1:]...)
	cmd.Env = ociConfig.ProcessConfig.Env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error running the main process - %s\n", err)
		os.Exit(1)
	}
	defer cmd.Wait()

}
