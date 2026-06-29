# yze-go-pkgstd

A [`yze`](https://github.com/gomatic/yze) analyzer (group `go`, category `structure`) enforcing the **per-package** standards of the gomatic three-tier CLI layout. For a command package (`internal/app/commands/<cmd>/`) it checks:

- the first declaration is the **const block** (`name`, `usage`, `argUsage`, `description`);
- a **`Command()`** entry point exists;
- the **domain** package is imported under the **`domain`** alias.

Cross-package correspondence (every command has a matching `internal/domain/<cmd>`, and vice versa) is enforced separately by the layout analyzer.

- **Rule:** `yze/go/pkgstd`
- **Binary:** `cmd/yze-go-pkgstd` runs it standalone.

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.
