package main

import "fmt"

type project struct {
	disable bool
	name    string
}

func main() {
	project1 := project{
		disable: false,
		name:    "project_01",
	}
	project2 := project{
		disable: false,
		name:    "project_02",
	}
	project3 := project{
		disable: true,
		name:    "project_03",
	}
	project4 := project{
		disable: false,
		name:    "project_04",
	}
	project5 := project{
		disable: true,
		name:    "project_05",
	}

	projects := []project{project1, project2, project3, project4, project5}
	fmt.Println("len(projects):", len(projects))
	fmt.Println("projects", projects[len(projects)-1])
	// for i, p := range projects {
	// 	if !p.disable {
	// 		projects[i] = projects[len(projects)-1]
	// 		projects = projects[:len(projects)-1]
	// 	}
	// }
	// fmt.Println("projects:", projects)
}
