# My Fyne Application #

This is a template for a Fyne GUI application.

## Environment Variables ##

The application can read certain values from the environment. Setting these variables can be done by the usual methods for the OS on which the application is running.

The variables that can be set are:

**FYNE_THEME**: This specifies whether to override the default OS theme with either "dark" or "light" theme variants.

## Folder Structure ##

The source code tries to follow the standard Go structure for laying out source code. More information on that structure can be found here [Golang Standards -- Project Layout](https://github.com/golang-standards/project-layout).

**internal**: Contains internal packages that are used by the application, but not intended to be exported as standalone packages.

**assets**: Assets for the application such as PNG files, etc.

## Bundling Assets ##

Go allows for the bundling of assets like images into the binary itself. This can be accomplished with the command:

```
fyne bundle --output <filename>.go <imagefile.ext>
```

For this project, an example might be:

```
fyne bundle --package ui --output internal/ui/bundle.go assets/Icon.png
```

This process has been automated using **go:generate** headers in main.go for this application. To regenerate the bundle.go file if an asset changes, use:

```
go generate
```

from the project root folder. Note that assets can be added to the bundle with the --append parameter:

```
fyne bundle --package ui --output internal/ui/bundle.go --append anotherimage.png
```

## Packaging for Desktop (MacOS) ##

To package the application in an app bundle named **MyFyneApplication** using **Icon.png** in the assets folder for the icon, run the following from the compiled output folder:

```
fyne package --os darwin --icon ../assets/Icon.png --name MyFyneApplication
```

Or if the code is already compiled (such as with fyne-cross) one can use:

```
fyne package --os darwin --exe fyne-cross/bin/darwin-arm64/myfyneapplication
```

## Cross Compiling ##

Cross compiling is accomplished by using **fyne-cross** (this tool requires docker). To cross compile, the following command was used:

Compiling **on a Mac** for a Windows machine (both arm64 and amd64), from the root folder of the code run:
```
fyne-cross windows -arch amd64,arm64 -icon assets/Icon.png
```

Compiling **on a Mac** for a Mac machine (both arm64 and amd64), from the root folder of the code run:
```
fyne-cross darwin -arch amd64,arm64 -icon assets/Icon.png
```

By default the compiled output will be in the **fyne-cross** folder, in sub-folders for each targeted OS and architecture. Use the builds in these folders for packaging (see above).
