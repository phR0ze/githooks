# githooks
Set of githooks to increment versions, update copyrights etc...

### Table of Contents
* [Add hooks to your project](#add-hooks-to-your-project)

## Add hooks to your project <a name="add-hooks-to-your-project"/></a>
```bash
# Clone githooks project
git clone https://github.com/phR0ze/githooks

# Create location in your project for hooks
mkdir ~/Projects/cyberlinux/.githooks

# Install githooks to project location
cp githooks/hooks/* ~/Projects/cyberlinux/.githooks

# Enable git hooks
cd ~/Projects/cyberlinux
git config core.hooksPath .githooks
```
