https://github.com/astral-sh/ruff

python linter

我一般用作 commit hook

示例: 

.git/hooks/pre-commit 可以在 commit 前执行，返回 0 才能 commit 成功

比如我经常写成下面这样, 实现如下效果：只有在 lint 之后才准许 commit:

```
#!/bin/bash

set -euxo pipefail
ruff check . --ignore E501
```