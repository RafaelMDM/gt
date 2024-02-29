# GT

## About

This program stores location bookmarks in a json file in the the user's config directory
e.g. `~/.config` on Linux.

## Installation

Make sure your [GOBIN](https://pkg.go.dev/cmd/go#hdr-Environment_variables) environment variable is set.

```bash
git clone https://github.com/RafaelMDM/gt.git && cd gt/cmd/gt
go install
```

Then add this function to your .bashrc or equivalent file:

```bash
function gt {
    go_bin=$(go env GOBIN)
    out=$("$go_bin/gt" "$@")

    if [ "$#" -eq 1 ]; then
        cd "$out"
    else
        echo "$out"
    fi
}
```

Since a child process can't change the cwd of its parent, this bash function will retrieve the path and cd to it.

## Usage

To add a new location bookmark:

```bash
gt add [name] [path]
# OR
gt [name] [path]

# Example: Save the cwd as "foo"
gt foo .
```

To remove a location bookmark:

```bash
gt rm [name]

# Example: Remove the "foo" bookmark:
gt rm foo
```

To list all locations:

```bash
gt list
# OR
gt
```

To go to a saved location bookmark:
```bash
gt [name]

# Example: Go to "foo"
gt foo
```

## Author

[Rafael Marques](https://github.com/RafaelMDM)
