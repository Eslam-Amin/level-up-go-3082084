package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`

	AlbumCount, SongCount int
}

// A PlaylistHeap is a max-heap of playlistEntries.
type PlaylistHeap []Song

func (ph PlaylistHeap) Len() int {
	return len(ph)
}

func (ph PlaylistHeap) Less(i, j int) bool {
	return ph[i].PlayCount > ph[j].PlayCount
}

func (ph PlaylistHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}

func (ph *PlaylistHeap) Push(x any) {
	*ph = append(*ph, x.(Song))
}

func (ph *PlaylistHeap) Pop() any {
	original := *ph
	ln := len(original)
	lastElement := original[ln-1]
	*ph = original[0 : ln-1]
	return lastElement
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	var playlist []Song
	pHeap := &PlaylistHeap{}
	if len(albums) == 0 {
		return playlist
	}

	heap.Init(pHeap)
	for i, f := range albums {
		firstSong := f[0]
		firstSong.AlbumCount, firstSong.SongCount = i, 0
		heap.Push(pHeap, firstSong)
	}
	for pHeap.Len() != 0 {
		p := heap.Pop(pHeap)
		song := p.(Song)
		playlist = append(playlist, song)
		if song.SongCount < len(albums[song.AlbumCount])-1 {
			nextSong := albums[song.AlbumCount][song.SongCount+1]
			nextSong.AlbumCount, nextSong.SongCount =
				song.AlbumCount, song.SongCount+1
			heap.Push(pHeap, nextSong)
		}
	}
	return playlist
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
