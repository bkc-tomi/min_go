# Golang CheetSheet

#### user.Current
現在のカレントユーザーを取得します。ユーザーディレクトリを参照可能。
しかし、クロスコンパイルの際はエラーになるかも？
>import "os/user"
>func Current() (*User, error)
>Current returns the current user.

#### filepath.Join
引数に与えられたパスをOSに合わせて一つのパスにつなげてくれる。
>import "path/filepath"
>func Join(elem ...string) string
>Join joins any number of path elements into a single path, separating them with an OS specific Separator. Empty elements are ignored. The result is Cleaned. However, if the argument list is empty or all its elements are empty, Join returns an empty string. On Windows, the result will only be a UNC path if the first non-empty element is a UNC path.

#### log.Fatal
Fatal は Print() の後に os.Exit(1) を呼び出すのと同じです。
>import "log"
>func Fatal(v ...interface{})
>Fatal is equivalent to Print() followed by a call to os.Exit(1).

#### os.MkdirAll
サブディレクトリも含めて作成したい時に使用する
>import "os"
>func MkdirAll(path string, perm FileMode) error
>MkdirAll creates a directory named path, along with any necessary parents, and returns nil, or else returns an error. The permission bits perm (before umask) are used for all directories that MkdirAll creates. If path is already a directory, MkdirAll does nothing and returns nil.

#### os.GetWd
現在のディレクトリのパスとエラーを返す
>import "os"
>func Getwd() (dir string, err error)
>Getwd returns a rooted path name corresponding to the current directory.

#### path.Match

>import "path"
>func Match(pattern, name string) (matched bool, err error)
>Match reports whether name matches the shell file name pattern. The pattern syntax is

#### http.HandleFunc
>import "net/http"
>func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
>HandleFunc registers the handler function for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.
#### http.ListenAndServe
>import "net/http"
>func ListenAndServe(addr string, handler Handler) error
>ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
>The handler is typically nil, in which case the DefaultServeMux is used.
以下のサイトを参照されたし
[GoでHTTPサーバー入門 Handler, HandleFunc, ServeMux](https://noumenon-th.net/programming/2019/09/12/handler/)

#### http.ResponseWriter
インターフェース
>import "net/http"
>A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
#### http.Request
HTTPリクエスト型
>import "net/http"
>A Request represents an HTTP request received by a server or to be sent by a client.
#### http.NotFound
該当するページがない場合に４０４エラーを返す。
>import "net/http"
>func NotFound(w ResponseWriter, r *Request)
>NotFound replies to the request with an HTTP 404 not found error.
#### http.Request.FormFile
>import "net/http"
>func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
>FormFile returns the first file for the provided form key. FormFile calls ParseMultipartForm and ParseForm if necessary.
#### http.Error
リクエストがエラーに該当する場合そのエラーメッセージとHTTPコードをクライアントに返す。
>import "net/http"
>func Error(w ResponseWriter, error string, code int)
>Error replies to the request with the specified error message and HTTP code. It does not otherwise end the request; the caller should ensure no further writes are done to w. The error message should be plain text.
#### http.StatusText
>import "net/http"
#### http.StatusInternalServerError
>import "net/http"
#### http.Redirect
リダイレクトを行う。
>import "net/http"
>func Redirect(w ResponseWriter, r *Request, url string, code int)
>Redirect replies to the request with a redirect to url, which may be a path relative to the request path.

#### io.Copy
srcをdstにコピーする。
>import "io"
>func Copy(dst Writer, src Reader) (written int64, err error)
>Copy copies from src to dst until either EOF is reached on src or an error occurs. It returns the number of bytes copied and the first error encountered while copying, if any.

#### filepath.Base
パスの最後の要素を取り出す。つまりほとんどの場合ファイル名。
>import "path/filepath"
>func Base(path string) string
>Base returns the last element of path. Trailing path separators are removed before extracting the last element. If the path is empty, Base returns ".". If the path consists entirely of separators, Base returns a single separator.

#### os.RemoveAll
指定されたパスのディレクトリと中のファイルを全て消去する。
>import "os"
>func RemoveAll(path string) error
>RemoveAll removes path and any children it contains. It removes everything it can but returns the first error it encounters. If the path does not exist, RemoveAll returns nil (no error). If there is an error, it will be of type *PathError.

#### os.Create
引数にパスも含めてファイル名を指定し、作成する。
>import "os"
>func Create(name string) (*File, error)
>Create creates or truncates the named file. If the file already exists, it is truncated. If the file does not exist, it is created with mode 0666 (before umask). If successful, methods on the returned File can be used for I/O; the associated file descriptor has mode O_RDWR. If there is an error, it will be of type *PathE
