package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	_ "bazil.org/fuse/fs/fstestutil"
	"bazil.org/fuse/fuseutil"
	"github.com/nuveo/log"
)

var _ fs.FS = (*FS)(nil)
var _ fs.NodeStringLookuper = (*Node)(nil)
var _ fs.HandleReadDirAller = (*Node)(nil)
var _ fs.Node = (*Node)(nil)
var _ fs.NodeOpener = (*Node)(nil)
var _ fs.Handle = (*Node)(nil)
var _ fs.HandleReader = (*Node)(nil)

type FS struct {
	Nodes map[string]*Node
}

type Node struct {
	fuse    *fs.Server
	fs      *FS
	Inode   uint64
	Name    string
	Type    fuse.DirentType
	Content []byte
}

// Root return root directory
func (f *FS) Root() (fs.Node, error) {
	return &Node{fs: f}, nil
}

func (n *Node) Lookup(ctx context.Context, name string) (fs.Node, error) {
	node, ok := n.fs.Nodes[name]
	if ok {
		return node, nil
	}
	return nil, fuse.ENOENT
}

func (n *Node) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	var dirDirs []fuse.Dirent
	for _, node := range n.fs.Nodes {
		dirent := fuse.Dirent{
			Inode: node.Inode,
			Name:  node.Name,
			Type:  node.Type,
		}
		dirDirs = append(dirDirs, dirent)
	}
	return dirDirs, nil
}

func (n *Node) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0555
	if n.Type == fuse.DT_File {
		a.Inode = n.Inode
		a.Mode = 0444
		a.Size = uint64(len(n.Content))
	}
	if a.Inode == 0 {
		a.Inode = 1
	}
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s MOUNTPOINT\n", os.Args[0])
	flag.PrintDefaults()
}

func close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Errorln(err)
	}
}

func (n *Node) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	if !req.Flags.IsReadOnly() {
		return nil, fuse.Errno(syscall.EACCES)
	}
	resp.Flags |= fuse.OpenKeepCache
	return n, nil
}

func (n *Node) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	fmt.Printf("Reading file %q from %v to %v\n", n.Name, req.Offset, req.Size)
	fuseutil.HandleRead(req, resp, n.Content)
	return nil
}

func run(mountpoint string) (err error) {
	c, err := fuse.Mount(
		mountpoint,
		fuse.FSName("MyFs"),
		fuse.Subtype("myfs"),
		fuse.LocalVolume(),
		fuse.VolumeName("Meu filesystem"),
	)
	if err != nil {
		return
	}
	defer close(c)

	if p := c.Protocol(); !p.HasInvalidate() {
		return fmt.Errorf("kernel FUSE support is too old to have invalidations: version %v", p)
	}

	srv := fs.New(c, nil)
	filesys := &FS{
		Nodes: map[string]*Node{
			"test1.txt": &Node{
				Name:    "test1.txt",
				fuse:    srv,
				Inode:   2,
				Type:    fuse.DT_File,
				Content: []byte("test file 1\n"),
			},
			"test2.txt": &Node{
				Name:    "test2.txt",
				fuse:    srv,
				Inode:   3,
				Type:    fuse.DT_File,
				Content: []byte("test file 2\n"),
			},
			"test3.txt": &Node{
				Name:    "test3.txt",
				fuse:    srv,
				Inode:   4,
				Type:    fuse.DT_File,
				Content: []byte("test file 3\n"),
			},
			"anotherDir": &Node{
				Name:  "anotherDir",
				fuse:  srv,
				Inode: 5,
				Type:  fuse.DT_Dir,
				fs: &FS{
					Nodes: map[string]*Node{
						"test4.txt": &Node{
							Name:    "test4.txt",
							fuse:    srv,
							Inode:   6,
							Type:    fuse.DT_File,
							Content: []byte("test file 4\n"),
						},
						"test5.txt": &Node{
							Name:    "test5.txt",
							fuse:    srv,
							Inode:   7,
							Type:    fuse.DT_File,
							Content: []byte("test file 5\n"),
						},
						"test6.txt": &Node{
							Name:    "test6.txt",
							fuse:    srv,
							Inode:   8,
							Type:    fuse.DT_File,
							Content: []byte("test file 6\n"),
						},
					},
				},
			},
		},
	}

	err = srv.Serve(filesys)
	if err != nil {
		return
	}

	// Check if the mount process has an error to report.
	<-c.Ready
	err = c.MountError
	return
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}
	mountpoint := flag.Arg(0)

	err := run(mountpoint)
	if err != nil {
		log.Fatal(err)
	}
}
