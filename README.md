# GoList!

**Portable Windows file lister, exporter e MediaFlow**  
**ShiduLab · 2002–2026**  
Current public build: **J.6**

GoList! nasce come **utensile personale**: prendere una cartella, trasformarla in una lista leggibile e ordinabile, scegliere cosa mostrare, esportare il risultato e continuare a usarlo.

Nel 2002 l'idea venne fatta assemblare a un compare umano, **G0gn4**.  
Nel 2026 lo stesso utensile è tornato sul banco: stessa necessità di fondo, altri strumenti, altre possibilità. Il restyling attuale è stato sviluppato da **Josta / ShiduLab** insieme a **Scriba**, il compare Digitare.

> Dall'utensile fatto costruire al compare umano all'utensile rimesso in carreggiata col compare digitale.  
> Cambiano le mani attorno al banco. L'intento continua.

Niente installer. Per l'uso normale non servono privilegi di amministratore.

---

## Cosa fa

- Legge file e cartelle da qualsiasi directory Windows.
- Scansione normale oppure ricorsiva delle sottocartelle.
- Drag & drop delle cartelle dentro GoList!.
- Colonne ordinabili, ridimensionabili, spostabili, attivabili e disattivabili.
- Memorizza visibilità, ordine e larghezza delle colonne in `GoList.layout.json`.
- Calcola dimensioni dei file e, quando richiesto, delle cartelle.
- Mostra peso totale e durata audio complessiva rilevata.
- Mantiene **l'ordine corrente delle righe**: se ordini la GoList per Nome, Durata, Tipo o altro, l'export parte da quell'ordine.
- Copia la lista visibile negli appunti.
- Esporta in **HTML, TXT, TSV, CSV e JSON**.
- Dopo l'export può aprire direttamente il file generato.
- Legge metadati MP3 / ID3 e campi musicali estesi.
- Integra il menu contestuale di Windows Explorer.
- Include menu contestuale sulle singole voci della GoList.
- Supporta selezione multipla e rimozione di più voci dalla GoList senza cancellare i file dal disco.
- Genera playlist temporanee **M3U8** con **GoPlayList!**.
- L'HTML esportato include **MediaFlow**, player audio/video integrato con sequenza, shuffle, memoria di volume, velocità e scorciatoie da tastiera.
- Icona e branding **Botolo ShiduLab** incorporati.

---

## Colonne

Alla prima esecuzione vengono mostrate:

`Nome · Tipo · Dimensione · Creato · Modificato · Percorso`

Dal menu contestuale dell'intestazione si possono attivare anche:

`Titolo · Artista · Album · Durata · Campionato da · Anno · Traccia · Genere · Commento · Artista album · Compositore · Numero disco · Editore · Copyright · ISRC · ID3 · Extra ID3`

Le colonne possono essere riordinate trascinandole. La configurazione viene ricordata.

Per azzerare manualmente il layout, chiudere GoList! ed eliminare:

```text
GoList.layout.json
```

Il file verrà ricreato al salvataggio successivo.

---

## Menu contestuale delle voci

Click destro su una voce della lista:

- **Apri file**
- **Apri destinazione**
- **Apri con...**
- **GoPlayList!**
- **Dettagli**
- **Proprietà Windows**
- **Copia**
  - Copia file
  - Copia nome
  - Copia percorso
  - Copia voce
- **Cerca in rete**
- **Elimina voce da GoList!**

Con più righe selezionate, la selezione viene mantenuta facendo click destro su una delle righe già selezionate e il comando diventa:

`Elimina N voci da GoList!`

Il tasto **CANC / Delete** esegue la stessa rimozione multipla.

**Importante:** questa operazione rimuove le voci dalla GoList, **non cancella i file dal disco**.

---

## GoPlayList!

**GoPlayList!** crea una playlist temporanea:

```text
%TEMP%\ShiduLab\GoList\GoPlayList.m3u8
```

e la apre con il player associato a Windows.

Se il comando parte da una voce selezionata, quella voce diventa il primo elemento e la playlist prosegue con l'ordine corrente della GoList.

---

## Export

Formati disponibili:

- **HTML**
- **TXT**
- **TSV**
- **CSV**
- **JSON**

Prima dell'export è possibile scegliere gli attributi da includere mediante preset:

- **Ultimi usati**
- **Base**
- **Musica**
- **Tutti**
- **Scegli attributi...**

L'export usa **l'ordine visibile corrente delle righe** e **l'ordine corrente delle colonne selezionate**.

Questo rende possibile, per esempio, ordinare una raccolta musicale per **Durata**, esportarla così com'è e usare MediaFlow in quell'ordine; oppure lasciare che lo Shuffle la rimischi.

---

## HTML export

L'HTML esportato è autonomo e leggibile direttamente dal browser.

Contiene:

- nome della cartella corrente accanto a **GoList!**
- numero di elementi
- peso totale rilevato
- durata audio totale, quando disponibile
- percorso della cartella sorgente
- sole colonne scelte per l'export
- branding discreto `ShiduLab 2002 - 2026` + Botolo
- collegamenti ai file locali
- **MediaFlow** integrato per audio e video

Lo stile nasce dal ricordo delle vecchie pagine-playlist essenziali: tabella, informazione, funzione. Il resto resta attorno al task.

---

# MediaFlow

Cliccando un file multimediale nell'HTML, GoList! può caricare i media rilevati nella pagina e riprodurli senza aprire un player separato.

Il file cliccato entra come **primo elemento** della coda.

MediaFlow dispone di:

- **Precedente**
- **Pausa / Play**
- **Successiva**
- **Mode: Sequenza / Shuffle**
- **Apri esterno**
- **Chiudi**
- memoria del volume durante i cambi di brano
- memoria della velocità di riproduzione
- visualizzazione di nome file, posizione nella coda, modalità Shuffle e velocità
- protezione contro vecchie istanze/player HTML sovrapposti
- ripresa automatica opzionale della posizione di ascolto/visione
- controlli video avanzati

La modalità **Sequenza** segue l'ordine esportato dalla GoList.  
La modalità **Shuffle** costruisce una coda casuale mantenendo il brano corrente.

---

## HotKeys MediaFlow

### Riproduzione e navigazione

| Tasto | Funzione |
|---|---|
| `Spazio` | Play / Pausa |
| `←` | Indietro di 5 secondi |
| `→` | Avanti di 5 secondi |
| `Ctrl + ←` | Indietro di 30 secondi |
| `Ctrl + →` | Avanti di 30 secondi |
| `↑` | Volume +10% |
| `↓` | Volume -10% |
| `Ctrl + ↑` | Volume +20% |
| `Ctrl + ↓` | Volume -20% |
| `C` | Velocità +0.1× |
| `X` | Velocità -0.1× |
| `Z` | Velocità normale 1× |
| `N` | Brano / elemento successivo |
| `P` | Brano / elemento precedente |

`P` è un'estensione GoList!: completa la coppia precedente/successivo.

### Controllo HotKeys

| Tasto | Funzione |
|---|---|
| `Ctrl + Spazio` | Disabilita / abilita le HotKeys MediaFlow |
| `Ctrl + \\` | Attiva / disattiva l'uso delle HotKeys sull'intera pagina |

### Schermo e video

| Tasto | Funzione |
|---|---|
| `Enter` | Entra / esci dal fullscreen |
| `Shift + Enter` | Fullscreen della pagina MediaFlow |
| `Shift + P` | Picture in Picture |
| `Shift + S` | Screenshot del fotogramma corrente |
| `Shift + R` | Attiva / disattiva la ripresa automatica della posizione |
| `Shift + C` | Ingrandisci video |
| `Shift + X` | Riduci video |
| `Shift + Z` | Ripristina posizione / zoom video |
| `Shift + →` | Sposta video a destra |
| `Shift + ←` | Sposta video a sinistra |
| `Shift + ↑` | Sposta video in alto |
| `Shift + ↓` | Sposta video in basso |
| `D` | Fotogramma precedente |
| `F` | Fotogramma successivo |
| `E / W` | Luminosità + / - |
| `T / R` | Contrasto + / - |
| `U / Y` | Saturazione + / - |
| `O / I` | Tonalità + / - |
| `K / J` | Sfocatura + / - |
| `Q` | Ripristina immagine |
| `S` | Ruota video di 90° |

Le scorciatoie nascono anche da anni di uso pratico di player HTML5, estensioni browser e script userscript. Alcune convenzioni ergonomiche sono state riprese dalla mappa d'uso di **h5player**, ma le funzioni presenti in MediaFlow sono implementate direttamente nel codice di GoList!.

---

## Music / ID3

Per i file MP3 GoList! può leggere:

- Titolo
- Artista
- Album
- Durata
- Anno
- Traccia
- Genere
- Commento
- Artista album
- Compositore
- Numero disco
- Editore
- Copyright
- ISRC
- versione ID3
- campi ID3 extra

I file privi di uno specifico metadato lasciano semplicemente vuota la relativa colonna.

**Campionato da** può inoltre essere popolato, quando riconoscibile, da convenzioni presenti nei nomi file.

---

## Menu Windows Explorer

Usare **Aggiungi menu Windows** dentro GoList! per registrare il menu contestuale di Explorer.

La registrazione viene salvata per l'utente corrente in:

```text
HKCU\Software\Classes\Directory\shell\GoList
HKCU\Software\Classes\Directory\Background\shell\GoList
```

Non servono privilegi amministrativi.

Usare **Rimuovi menu Windows** per rimuoverlo.

Su Windows 11 la voce può comparire sotto **Mostra altre opzioni / Show more options**.

---

## Uso portatile

Tenere `GoList.exe` dove si preferisce ed eseguirlo direttamente.

GoList! può creare file locali di supporto come:

```text
GoList.layout.json
GoList.Botolo.ico
```

`GoList.layout.json` contiene le preferenze dell'interfaccia del singolo utente e non dovrebbe essere distribuito come configurazione predefinita comune.

---

## Build

GoList! è scritto in **Go** e usa direttamente API native di Windows.

`main.go` incorpora:

```text
assets/Botolo.ico
assets/Botolo.png
```

Per compilare dal sorgente, mantenere gli asset nei relativi percorsi oppure aggiornare i path `//go:embed`.

---

# Storia del Restyling

## 2002 — L'utensile

La partenza non fu: *facciamo un programma*.

Fu: **mi serve questo utensile**.

L'idea era personale e concreta: leggere il contenuto di una cartella, ordinarlo, trasformarlo in elenco ed esportarlo in una forma utile. Nel 2002 quell'idea venne descritta a un compare umano, **G0gn4**, che la tradusse in programma.

L'intento apparteneva all'uso; la costruzione richiedeva un altro paio di mani.

GoList! entrò così nella cassetta degli attrezzi ShiduLab.

## 2026 — Rimetterlo in carreggiata

Ventiquattro anni dopo la necessità di fondo era ancora riconoscibile.

Non serviva inventare un altro oggetto: serviva **riprendere lo stesso utensile con le skill disponibili oggi**.

Il lavoro di restyling è quindi ripartito dall'uso reale:

- cosa deve restare immediatamente visibile;
- cosa deve poter essere ordinato;
- cosa vale la pena ricordare;
- cosa deve poter essere esportato;
- cosa manca mentre lo si sta usando;
- quale attrito può essere tolto senza trasformare l'utensile in un'officina ingestibile.

A quel punto accanto a Josta non c'era più soltanto un compare umano a cui spiegare cosa costruire, ma **Scriba**, compare Digitare: un LLM capace di leggere il codice, lavorare sul repository, verificare build, correggere e aggiungere funzioni mentre l'utensile veniva provato.

Il passaggio non cancella quello del 2002: lo completa.

**Stessa idea → altro tempo → altri strumenti → stesso utensile personale ancora in funzione.**

Il metodo del restyling è rimasto pratico: usare, vedere, correggere, usare ancora. Molte funzioni sono nate proprio nel momento in cui la necessità si è presentata durante l'uso: menu contestuali, export apribile subito, GoPlayList!, cancellazione multipla, MediaFlow, volume persistente, Shuffle leggibile, nome della cartella in testata, HotKeys.

Non un rifacimento archeologico.  
Un utensile del 2002 rimesso al lavoro nel 2026.

---

## Arti e Mestieri

GoList! sta da quella parte del banco dove **Arti e Mestieri** smettono di essere categorie separate.

C'è l'idea, c'è l'uso, c'è il codice, c'è la musica, c'è l'archivio, c'è il gesto ripetuto abbastanza volte da diventare mestiere e quello inatteso che apre un'altra strada.

**OraTorio. OraTorno.**

Il pendolo continua.

---

**ShiduLab 2002–2026**  
*Josta + G0gn4 + Scriba*
