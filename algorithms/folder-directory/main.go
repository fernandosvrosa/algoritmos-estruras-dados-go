package main

import (
	"fmt"
	"strings"
)

// Given a list of folders, print the path of a given folder from a root. Print "no-path" if there is no path from the root folder. Root-level folders will have 0 as the id. The structure of Folder is like this:
// class Folder {
//    int id;
//    int[] subfolders;
//    String name;
// }

// Ex:
// list = [
//     Folder(0, [7, 3], "abc"),
//     Folder(0, [4], "xyz"),
//     Folder(8, [9], "def"),
//     Folder(7, [9], "ijk"),
//     Folder(9, [], "lmn"),
//     Folder(4, [], "opq"),
// ]

// printPath(list, 9) = "abc" -> "ijk" -> "lmn"
// printPath(list, 8) = "no-path"
//
//

type Folder struct {
	id         int
	subfolders []int
	name       string
}

// Função que encontra o caminho de uma pasta específica até a raiz
func printPath(folders []Folder, targetID int) string {
	// Criar mapas para facilitar a navegação
	folderMap := make(map[int]Folder)
	parentMap := make(map[int]int)

	// Popular os mapas com dados das pastas
	for _, folder := range folders {
		folderMap[folder.id] = folder
		for _, subID := range folder.subfolders {
			parentMap[subID] = folder.id
		}
	}

	// Verificar se a pasta alvo está no mapa
	if _, exists := folderMap[targetID]; !exists {
		return "no-path"
	}

	// Construir o caminho da pasta alvo até a raiz
	var path []string
	currentID := targetID
	for {
		folder, exists := folderMap[currentID]
		if !exists {
			return "no-path"
		}
		path = append([]string{folder.name}, path...)
		parentID, hasParent := parentMap[currentID]
		if !hasParent {
			break
		}
		currentID = parentID
	}

	// Verificar se chegamos à raiz
	if _, isRoot := folderMap[currentID]; !isRoot {
		return "no-path"
	}

	// Retornar o caminho como string
	return strings.Join(path, " -> ")
}

func main() {
	// Lista de exemplo de pastas
	folders := []Folder{
		{id: 0, subfolders: []int{7, 3}, name: "abc"},
		{id: 0, subfolders: []int{4}, name: "xyz"},
		{id: 8, subfolders: []int{9}, name: "def"},
		{id: 7, subfolders: []int{9}, name: "ijk"},
		{id: 9, subfolders: []int{}, name: "lmn"},
		{id: 4, subfolders: []int{}, name: "opq"},
	}

	// Testar a função com diferentes IDs
	fmt.Println(printPath(folders, 9)) // Deve imprimir "abc -> ijk -> lmn"
	fmt.Println(printPath(folders, 8)) // Deve imprimir "no-path"
}
