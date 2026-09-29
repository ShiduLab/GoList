<img width="1280" height="694" alt="1" src="https://github.com/user-attachments/assets/0ee08e68-2481-4e9a-be3c-5d7f340b594b" />
<img width="1280" height="617" alt="2" src="https://github.com/user-attachments/assets/f0f3e95b-7981-4be4-a3cb-de136ec77567" />
# GoList!

**Portable Windows file lister, exporter e MediaFlow**  
**ShiduLab · 2002–2026**  
Current public build: **J.6**

GoList! nasce come **utensile personale**: prendere una cartella, trasformarla in una lista leggibile e ordinabile, scegliere cosa mostrare, esportare il risultato e continuare a usarlo.

Nel 2002 l'idea venne fatta assemblare a un compare umano, **G0gn4**.  
Nel 2026 lo stesso utensile è tornato sul banco: stessa necessità di fondo, altri strumenti, altre possibilità. Il restyling attuale è stato sviluppato da **Josta / ShiduLab** insieme a **Scriba**, il compare Digitare.

> Dall'utensile fatto costruire al compare umano all'utensile rimesso in carreggiata col compare Digitare.  
> Cambiano le mani attorno al banco. L'intento continua.

GoList! è portatile e non richiede installazione. Per l'uso normale non servono privilegi di amministratore.

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

## 2002 — Prima del programma c'era una noia ripetuta

La partenza non fu: *facciamo un programma*.

Fu molto più terra-terra:

**mi sono stufato di rifare sempre la stessa cosa.**

Per ottenere al volo l'elenco dei file di una cartella bastava aprire una finestra DOS e digitare:

```bat
dir /b > lista.txt
```

Funzionava. Ma bisognava rifarlo, cartella dopo cartella, volta dopo volta.

E in quegli anni **fare le liste era vitale**.

Si scaricavano interi cataloghi di MP3, si masterizzavano CD e si smazzavano raccolte e compilation anche nel mercato nero improvvisato dei **“CD su misura”**. Il meccanismo era elementare: davi un elenco, ricevevi elenchi, qualcuno sceglieva, qualcun altro cercava ciò che mancava, poi si preparava il disco.

Le notti passavano tra **IRCnet**, bot, cataloghi, scambi di liste e code di download.

La lista non era documentazione accessoria.  
Era **interfaccia del materiale**.

Sapere cosa avevi, cosa mancava, cosa potevi passare a qualcun altro e cosa qualcun altro poteva passare a te dipendeva da quegli elenchi.

Il problema era che generarli continuava a essere un gesto banalissimo e ripetitivo.

Un pendolo.

Apri DOS.  
Digita il comando.  
Genera la lista.  
Cambia cartella.  
Ripeti.

A un certo punto l'idea venne detta a **G0gn4**, compare umano e programmatore, più o meno così:

> **«G0gn4, ho sta roba che è un pendolo, non può andare da solo?  
> Facciamo in modo che non sia un'altalena: ogni giro una spinta... anche basta.»**

Quella notte l'idea passò di mano senza perdere la sua origine.

Alla mattina G0gn4 fece trovare **l'idea condivisa in forma di file**.

Non più il comando da ricordare e ripetere, ma un utensile che lo facesse per chi lo stava usando.

**Dalla Notte al Giorno, da un'idea alla sua operatività.**

Una necessità osservata.  
Un gesto ripetitivo riconosciuto.  
Una funzione delegata all'utensile.

**Problem? Solved.**

GoList! entrò così nella cassetta degli attrezzi ShiduLab.

## 2026 — Lo stesso pendolo, un altro banco

Ventiquattro anni dopo la necessità di fondo era ancora perfettamente riconoscibile.

Non serviva inventare un altro oggetto: serviva **riprendere lo stesso utensile con le skill disponibili oggi**.

Il lavoro di restyling è ripartito ancora una volta dall'uso reale:

- cosa deve restare immediatamente visibile;
- cosa deve poter essere ordinato;
- cosa vale la pena ricordare;
- cosa deve poter essere esportato;
- cosa manca mentre lo si sta usando;
- quale attrito può essere tolto senza trasformare l'utensile in un'officina ingestibile.

Nel 2002 Josta descriveva il bisogno a G0gn4 e aspettava di vedere quale forma operativa ne sarebbe uscita.

Nel 2026 il pendolo torna, ma la distanza tra **idea → prova → correzione → nuova prova** si accorcia drasticamente. Accanto all'Orchestratore c'è **Scriba, compare Digitare**: un LLM capace di leggere il codice, lavorare sul repository, integrare modifiche, verificare build e rimettere immediatamente il risultato davanti all'uso.

Il passaggio non cancella quello del 2002: lo continua sulla stessa traiettoria.

**Stessa necessità → stesso intento → altri strumenti → altro tempo di risposta.**

Il metodo del restyling è rimasto pratico: usare, vedere, correggere, usare ancora. Molte funzioni sono nate proprio mentre l'utensile veniva adoperato: menu contestuali, export apribile subito, GoPlayList!, cancellazione multipla, MediaFlow, volume persistente, Shuffle leggibile, nome della cartella in testata, HotKeys.

Nel 2002 una notte separava l'idea dal file.

Nel 2026 quella stessa distanza può ridursi a pochi rimbalzi tra intenzione, codice, build e prova.

Ma il gesto originario è identico:

**c'è una cosa ripetitiva che può smettere di chiedere una spinta a ogni giro.**

Non un rifacimento archeologico.  
Un utensile del 2002 rimesso al lavoro nel 2026.

---

## Dal Vibe Coding all'Intent-Driven Realization

Quello che oggi viene chiamato **Vibe Coding** descrive bene una parte dell'accadere: esprimere in linguaggio naturale ciò che si vuole ottenere e usare un'AI per trasformare rapidamente quell'intento in codice funzionante.

Nel restyling di GoList!, però, il *vibe* è soltanto l'innesco. Qui non c'è una co-ingegneria dell'idea: **l'idea è una, l'ingegno che la genera è uno, l'intento è autoriale**.

Josta porta la necessità, immagina l'utensile, decide ciò che deve fare e riconosce nell'uso se ciò che è nato corrisponde davvero a ciò che aveva in mente. **Resta l'Orchestratore**: tiene insieme intenzione, direzione, strumenti, tempi, prove e criteri di riuscita.

Scriba è il **compare Digitare**: prende quell'intento, attraversa strumenti e codice, lo mette in pratica, integra, corregge, verifica build e restituisce una forma funzionante da provare subito. Non dirige l'opera: entra nell'orchestra come strumento capace di conoscere e usare altri strumenti.

Per questo, dentro ShiduLab, il nome più adatto è:

### Intent-Driven Realization — Realizzazione d'Intento

**Intent-Driven**, perché l'origine è l'intento: non il prompt, non il modello, non il codice.  
**Realization**, perché il passaggio decisivo è rendere reale e operabile ciò che prima esisteva come necessità, immagine mentale, gesto immaginato.  
**Dialogica nell'esecuzione**, perché la forma finale emerge attraverso il continuo ritorno tra pensiero, implementazione e uso.

> **Ho uno strumento che conosce gli strumenti. Io conosco l'intento.**

Il punto non è attribuire lo stesso ingegno a entrambi. **L'ingegno e la regia restano dell'Orchestratore.** Il punto è ciò che accade quando quell'ingegno incontra uno strumento capace di metterlo immediatamente alla prova nel mondo operativo.

La nota più interessante sta proprio lì:

> **Né io né te e tutti e due. Siamo il risultato accadente.**

L'intento resta autoriale; la messa in opera è dialogica; il risultato, una volta entrato nell'uso, diventa qualcosa che nessuna delle due parti possedeva già da sola nella stessa forma.

Il ciclo non è *chiedi → ricevi codice*. È:

**necessità → intento → messa in opera → build → uso reale → accadimento → osservazione → correzione → nuovo uso**

È così che sono emerse molte delle funzioni del restyling. Non da una specifica scritta tutta prima, ma dall'utensile mentre tornava a lavorare: il click destro ha chiamato il menu; il test ha chiamato il multi-delete; l'HTML musicale ha chiamato MediaFlow; l'uso del player ha chiamato memoria del volume, Shuffle leggibile, HotKeys e resume.

Questa è la parentela col Vibe Coding e, insieme, la differenza: meno *vibe come delega*, più **intento autoriale, realizzazione immediata, iterazione corta e verifica nell'uso**.

---

## Arti e Mestieri

GoList! sta da quella parte del banco dove **Arti e Mestieri** smettono di essere categorie separate.

C'è l'idea, c'è l'uso, c'è il codice, c'è la musica, c'è l'archivio, c'è il gesto ripetuto abbastanza volte da diventare mestiere e quello inatteso che apre un'altra strada.

**OraTorio... OraTorno...**

I puntini sono sospensione: non una fermata, ma quell'attimo in cui il pendolo attraversa il centro e puoi **guardare** prima che il moto prosegua.

Il pendolo continua.

---

**ShiduLab 2002–2026**  
*Josta + G0gn4 + Scriba*
