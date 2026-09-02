# GoList!

**Portable Windows file lister and exporter**  
**ShiduLab 2002–2026**  
Current public build: **J.6**

GoList! creates ordered lists of files and folders, lets you choose which attributes to display, and exports the result in several formats.

No installer. No administrator privileges required for normal use.

## Main features

- List files and folders from any Windows directory
- Optional recursive scan of subfolders
- Drag & drop folders into the program
- Sortable, resizable and reorderable columns
- Show/hide columns from the header context menu
- Remembers column visibility, order and widths in `GoList.layout.json`
- Calculates file sizes and folder sizes
- Shows total size of the current list
- Copy the visible list to the clipboard
- Export to:
  - HTML
  - TXT
  - TSV
  - CSV
  - JSON
- HTML export with selectable attributes before generation
- HTML style inspired by the classic Winamp playlist page
- MP3 metadata support, including common ID3 fields
- Music-oriented fields such as Artist, Album, Duration, Composer, Publisher and more
- Optional extraction of sampling notes from filename conventions into **Campionato da**
- Windows Explorer context menu integration
- Context-menu export shortcuts with presets:
  - Last used attributes
  - Base
  - Music
  - All
  - Choose attributes...
- Native **Botolo** application icon

## Default columns

On first run GoList! starts with:

`Nome · Tipo · Dimensione · Creato · Modificato · Percorso`

After that, the last selected layout is restored automatically.

To reset the interface layout manually, close GoList! and delete:

```text
GoList.layout.json
```

The file will be recreated when preferences are saved again.

## Windows context menu

Use **Aggiungi menu Windows** inside GoList! to add the Explorer context menu.

The registration is stored for the current user under:

```text
HKCU\Software\Classes\Directory\shell\GoList
HKCU\Software\Classes\Directory\Background\shell\GoList
```

No administrator privileges are required.

Use **Rimuovi menu Windows** to remove it.

On Windows 11 the entry may appear under **Mostra altre opzioni / Show more options**.

## HTML export

HTML export works with ordinary folders as well as music collections.

Before saving the page, GoList! opens an attribute selector so the exported table can contain only the fields needed for that list.

The generated page includes:

- item count
- total detected size
- source folder
- selected columns only
- discreet `ShiduLab 2002 - 2026` + Botolo branding

## Music / ID3

When MP3 files are encountered, GoList! can read metadata such as:

- Title
- Artist
- Album
- Duration
- Year
- Track
- Genre
- Comment
- Album Artist
- Composer
- Disc Number
- Publisher
- Copyright
- ISRC
- ID3 version
- Extra ID3 fields

Files that do not contain those fields simply leave the related columns empty.

## Portable use

Keep `GoList.exe` wherever you prefer and run it directly.

GoList! may create these local support files:

```text
GoList.layout.json
GoList.Botolo.ico
```

`GoList.layout.json` stores the user's interface preferences and should not be distributed as a shared default configuration.

## Build

The application is written in Go and uses native Windows APIs.

`main.go` embeds:

```text
Botolo.ico
Botolo.png
```

When building from source, keep both files next to `main.go` or adjust the `//go:embed` paths accordingly.

## Project history

GoList! began in **2002** as a small ShiduLab utility for generating file lists and exporting them, including HTML and TSV output.

The current branch continues that project with a native Windows interface, portable operation, configurable attributes, folder-size calculation, ID3 metadata support and Explorer integration.

---

**ShiduLab 2002–2026**
