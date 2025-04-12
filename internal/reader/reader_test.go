package reader

import (
	"os"
	"path/filepath"
	"testing"
)

// mockStdinReader returns a *stdinReader that reads from a mocked pipe.
func mockStdinReader(input string) (*stdinReader, func(), error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	if _, err := w.WriteString(input); err != nil {
		return nil, nil, err
	}
	w.Close()

	reader := &stdinReader{source: r}
	cleanup := func() { r.Close() }
	return reader, cleanup, nil
}

// setupReader abstracts setting up either a mock or real stdinReader.
func setupReader(t *testing.T, input string, useRealStdin bool) (*stdinReader, func()) {
	t.Helper()

	if useRealStdin {
		return &stdinReader{source: os.Stdin}, func() {}
	}

	r, cleanup, err := mockStdinReader(input)
	if err != nil {
		t.Fatalf("failed to create mock stdin: %v", err)
	}
	return r, cleanup
}

func TestStdinReader_Read(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		useRealStdin bool
		want         string
		wantErr      bool
	}{
		{
			name:    "valid piped input",
			input:   "Hello, world!\n",
			want:    "Hello, world!",
			wantErr: false,
		},
		{
			name:    "empty piped input",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:         "non-piped input (real terminal)",
			useRealStdin: true,
			want:         "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.useRealStdin && os.Getenv("CI") == "true" {
				t.Skip("Skipping real stdin test in CI environment")
			}

			reader, cleanup := setupReader(t, tt.input, tt.useRealStdin)
			defer cleanup()

			got, err := reader.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Read() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFileReader_Read(t *testing.T) {
	t.Run("valid file input", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "testfile")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		content := "File content for testing."
		if _, err := tmpFile.WriteString(content); err != nil {
			t.Fatalf("failed to write to temp file: %v", err)
		}
		tmpFile.Close()

		reader, err := NewFileReader(tmpFile.Name())
		if err != nil {
			t.Fatalf("failed to create fileReader: %v", err)
		}

		got, err := reader.Read()
		if err != nil {
			t.Fatalf("unexpected read error: %v", err)
		}
		if got != content {
			t.Errorf("Read() = %q, want %q", got, content)
		}
	})

	t.Run("non-existent file", func(t *testing.T) {
		path := filepath.Join(os.TempDir(), "nonexistent_file.txt")

		reader, err := NewFileReader(path)
		if err != nil {
			t.Fatalf("unexpected error creating fileReader: %v", err)
		}

		_, err = reader.Read()
		if err == nil {
			t.Error("expected error when reading nonexistent file, got nil")
		}
	})

	t.Run("unreadable file", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "unreadable")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.Close()

		// Remove read permissions
		if err := os.Chmod(tmpFile.Name(), 0222); err != nil {
			t.Fatalf("failed to make file unreadable: %v", err)
		}
		defer os.Chmod(tmpFile.Name(), 0644) // Restore perms

		reader, err := NewFileReader(tmpFile.Name())
		if err != nil {
			t.Fatalf("failed to create fileReader: %v", err)
		}

		_, err = reader.Read()
		if err == nil {
			t.Error("expected error when reading unreadable file, got nil")
		}
	})
}
