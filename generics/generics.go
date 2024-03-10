package generics

import (
    "fmt"
)

// 引数の中に基本型が含まれているときはジェネリクスの使いどころな気がする。
// パフォーマンスに関しては、https://blog.stackademic.com/golang-design-generics-vs-interfaces-how-it-really-works-under-the-hood-c1e5057a7ae6 （後で読む）
func lookBack[memories photoPrint | digitalPhoto | string](mems []memories) {
    for _, mem := range mems {
        fmt.Println("I feel nostalgic looking at", mem)
    }
}

func GenericsMain() {
    var photoAlbum = []photoPrint{{"友達とのスキーの写真"}, {"旅行先で取った写真"}}
    lookBack(photoAlbum)
    var googlePhotos = []digitalPhoto{{"家族との登山の写真"}, {"食べ物の写真"}}
    lookBack(googlePhotos)
    var misc = []string{"a favorite childhood anime"}
    lookBack(misc)
}
