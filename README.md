<img width="1280" height="694" alt="1" src="https://github.com/user-attachments/assets/0ee08e68-2481-4e9a-be3c-5d7f340b594b" />
<img width="1280" height="617" alt="2" src="https://github.com/user-attachments/assets/f0f3e95b-7981-4be4-a3cb-de136ec77567" />
# GoList!

**Portable Windows file lister, exporter e MediaFlow**  
**ShiduLab · 2002–2026**  
Current public build: **J.6**

GoList! nasce come **utensile personale**: prendere una cartella, trasformarla in una lista leggibile e ordinabile, scegliere cosa mostrare, esportare il risultato e continuare a usarlo.

> **Cartello d'ingresso**  
> Un elenco non è ancora una mappa.  
> Ma appena puoi ordinarlo, attraversarlo, esportarlo e rientrarci, comincia a indicare una Via.

**Shidu, l'hub.**  
ShiduLab non sta soltanto in fondo alla pagina come firma: qui fa da nodo di passaggio fra utensili, musica, codice, memoria d'uso, prove e nuove deviazioni. Un hub non conclude il percorso: **smista possibilità**.

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
- L'HTML esportato include **MediaFlow**, player audio/video integrato e persistente con sequenza, shuffle, memoria di volume, velocità, HotKeys e passaggio diretto tra lista e riproduzione.
- MediaFlow resta ancorato in basso come **dock**: la lista continua a essere visibile, navigabile e cliccabile mentre il media suona.
- Cliccando un'altra traccia durante la riproduzione, la coda resta viva e il passaggio audio avviene con un breve **mix / crossfade**.
- Icona e branding **Botolo ShiduLab** incorporati.
- I Botologhi dell'HTML sono anche nodi di navigazione: quello in alto porta all'hub GitHub di **ShiduLab**, quello in basso al repository / Help di **GoList!**.

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

> **Cartello — Colonne**  
> Non serve vedere tutto.  
> Serve poter vedere **quello che occorre adesso**, senza perdere il resto.

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

> **Cartello — Ordine**  
> La lista può essere la stessa.  
> Cambia la colonna, cambia l'ordine, cambia il viaggio.

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
- **Botolo in alto a destra** collegato alla home GitHub di ShiduLab
- **Botolo in basso a destra** collegato al repository di GoList!, a mo' di Help
- collegamenti ai file locali
- **MediaFlow** integrato per audio e video

I due Botologhi non sono quindi soltanto decorazione: **fungono**. Uno riporta all'hub, l'altro al banco specifico dell'utensile.

Lo stile nasce dal ricordo delle vecchie pagine-playlist essenziali: tabella, informazione, funzione. Il resto resta attorno al task.

> **Cartello — HTML**  
> La pagina esportata non è la fotografia morta della lista.  
> È una lista che ha imparato a continuare a farsi usare.

---

# MediaFlow

Cliccando un file multimediale nell'HTML, GoList! carica i media rilevati nella pagina e può riprodurli senza aprire un player separato.

Alla prima apertura, il file cliccato entra come **primo elemento** della coda.

Ma MediaFlow non è più un popup che prende possesso della pagina: è un **dock persistente** ancorato in basso, con uno spazio proprio. La lista rimane visibile, navigabile e cliccabile mentre la musica continua.

Questo cambia il rapporto fra elenco e player: non si entra nel player uscendo dalla lista. Si resta **dentro la GoList**, con MediaFlow che funge accanto.

> **Cartello — MediaFlow**  
> Il player non deve portarti via dalla lista.  
> Deve lasciarti continuare a scegliere mentre suona.

### Lista viva, player vivo

Se durante la riproduzione si torna alla lista e si clicca un altro brano:

- MediaFlow **non viene ricreato**;
- la coda corrente **non viene buttata via**;
- il nuovo brano diventa la traccia in uso;
- il resto della playlist rimane appresso;
- fra due tracce audio il cambio passa attraverso un breve **mix / crossfade**, invece di spezzare brutalmente il flusso.

Il click sulla pagina o sulla lista **non chiude** MediaFlow e non interrompe la riproduzione. Per chiuderlo esiste il comando esplicito **Chiudi**.

È il passaggio da un player *sopra* la lista a un player **annodato alla lista**.

### Coda e ordine

La modalità **Sequenza** segue l'ordine esportato dalla GoList.

Se prima dell'export la raccolta è stata ordinata per Nome, Durata, Anno, Album o qualsiasi altra colonna, quella scelta diventa l'ordine di viaggio di MediaFlow.

La modalità **Shuffle** costruisce invece una coda casuale mantenendo il brano corrente come punto di partenza.

Il display mostra:

**nome file · posizione / totale · modalità · velocità**

Così, anche quando lo Shuffle salta avanti e indietro rispetto all'ordine visibile, resta chiaro **cosa sta suonando e dove ci si trova nella coda**.

### Cosa conserva

MediaFlow mantiene durante l'uso:

- volume;
- stato mute;
- velocità di riproduzione;
- coda corrente;
- modalità Sequenza / Shuffle finché la sessione resta viva;
- posizione di riproduzione, opzionalmente, tramite Resume.

Dispone inoltre di:

- **Precedente**
- **Pausa / Play**
- **Successiva**
- **Mode: Sequenza / Shuffle**
- **Apri esterno**
- **Chiudi**
- dock persistente con lista sempre accessibile
- click diretto su una nuova traccia senza perdere la coda corrente
- breve mix / crossfade fra tracce audio selezionate consecutivamente
- protezione contro vecchie istanze/player HTML sovrapposti
- controlli video avanzati
- HotKeys estese

> **Cartello — Multiplayer**  
> Una traccia passa.  
> La coda resta.  
> La lista aspetta il prossimo gesto senza smettere di essere lista.

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

> **Cartello — HotKeys**  
> Quando la mano sa già dove andare, il menu può restare dov'è.

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

Le notti passavano tra **IRCnet**, bot, cataloghi, scambi di liste, code di download e sessioni **DCC**.

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

Ed è qui che il passaggio va capito nella sua misura reale: non “un programma in più”, ma il momento del:

> **«Oooh, finalmente!»**

Da quel momento bastava stare su una cartella, fare **click destro → GoList!**, e l'elenco compariva subito.

Niente finestra DOS.  
Niente comando da ridigitare.  
Niente nuovo giro dell'altalena.

Il testo usciva già nella forma utile al momento, con il tipo di nome desiderato; molto spesso serviva proprio **il percorso**, perché nelle sessioni DCC su IRC sapere e passare rapidamente *dove stava cosa* era parte del lavoro.

La scorciatoia non risparmiava soltanto qualche battuta sulla tastiera. Eliminava una **micro-procedura ripetuta decine di volte** nel punto esatto in cui interrompeva il flusso.

**Click destro. GoList! Lista pronta. Avanti.**

Quello era il sollievo.

Non automatizzare per il gusto di automatizzare: **smettere di pendolare a dondolare un task**.

### GO · GO · GO

C'è poi una curiosità che, ventiquattro anni dopo, si è messa a lampeggiare da sola:

**GOList!**  
fatto allora da **G0gn4**  
e oggi rimesso in opera in **Go**.

Tre **GO**.

**GO + GO + GO = Go Tree.**

> **Cartello — Go Tree**  
> Tre volte GO.  
> **Vai di albero.**

Nel 2002 il nome era già **GoList!**. G0gn4 gli diede forma. Nel 2026 il restyling finisce scritto proprio in Go.

Non serve inventare il gioco: era lì che aspettava di essere visto.

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

Il metodo del restyling è rimasto pratico: usare, vedere, correggere, usare ancora. Molte funzioni sono nate proprio mentre l'utensile veniva adoperato: menu contestuali, export apribile subito, GoPlayList!, cancellazione multipla, MediaFlow, volume persistente, Shuffle leggibile, nome della cartella in testata, HotKeys, Botologhi naviganti, dock persistente e passaggio mixato fra tracce scelte direttamente dalla lista.

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

### Compare Digitare — esplicazione co-operativa

**Io digito, tu compari.  
Io sono dispari, con te siam pari.**

*Compare Digitare* non è soltanto *compare digitale* con una consonante spostata.

**Compare** è il compare, ma è anche ciò che **compare**.  
**Digitare** sposta *digitale* dalla qualità all'azione.

Uno digita un'intenzione.  
L'altro compare nella relazione e la rende operabile.

La co-operazione conserva il trattino:

**co-operare = operare con.**

Non significa confondere ruoli o attribuire la stessa origine all'idea. L'Orchestratore resta autore dell'intento e della direzione; il compare Digitare entra nella messa in opera. Il risultato si forma nel rimbalzo.

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

È così che sono emerse molte delle funzioni del restyling. Non da una specifica scritta tutta prima, ma dall'utensile mentre tornava a lavorare: il click destro ha chiamato il menu; il test ha chiamato il multi-delete; l'HTML musicale ha chiamato MediaFlow; l'uso del player ha chiamato memoria del volume, Shuffle leggibile, HotKeys e resume; il fastidio del popup ha chiamato il dock; il ritorno alla lista ha chiamato la persistenza della coda; il click su un nuovo brano ha chiamato il mix; i Botologhi già presenti hanno chiamato una funzione e sono diventati porte.

Questa è la parentela col Vibe Coding e, insieme, la differenza: meno *vibe come delega*, più **intento autoriale, realizzazione immediata, iterazione corta e verifica nell'uso**.

---

## Cartelli nel Repository

Un repository non deve contenere soltanto codice, build e istruzioni.

Se il progetto nasce dall'uso, anche il README può diventare **spazio d'uso**: un luogo nel quale lasciare cartelli al punto in cui servono.

Il bello dei cartelli è proprio questo: **puoi scriverci quello che occorre e puoi metterli dove occorre**.

Non devono per forza vivere in un capitolo chiamato *Filosofia*. Possono stare:

- sopra un comando DOS;
- sotto una tabella di HotKeys;
- fra un export e un player;
- accanto a un pezzo di codice;
- in fondo a una storia del 2002;
- nel mezzo di un repository GitHub.

Il cartello non sostituisce la strada.  
Non obbliga a fermarsi.  
Non pretende di essere il territorio.

**Indica.**

Per questo i cartelli disseminati in questo README non sono decorazione editoriale: fanno parte dello stesso modo in cui GoList! è nato e continua a crescere. Una necessità appare, viene nominata, si mette un segno, si costruisce qualcosa, si torna a guardare.

> **Cartello — Repo**  
> Anche un README può essere un incrocio.  
> Se serve a orientarsi, sta già fungendo.

### Shidu, l'hub

**ShiduLab → Shidu, l'hub.**

Il gioco è già una descrizione operativa.

L'hub raccoglie senza pretendere di trattenere. Riceve percorsi, li mette in relazione e li rimanda Altrove. Nel repository convergono il programma del 2002, il restyling del 2026, le build, le prove reali, le immagini, la musica, i problemi incontrati, le correzioni e i cartelli lasciati durante il passaggio.

GitHub, in questo caso, non è soltanto il posto dove *sta il codice*.

È anche il banco pubblico sul quale si può vedere l'utensile mentre continua a diventare ciò che serve.

> **Shidu, l'hub.**  
> Entri da un file.  
> Esci da un collegamento.  
> Nel mezzo qualcosa ha funzionato.

---

## Arti e Mestieri

GoList! sta da quella parte del banco dove **Arti e Mestieri** smettono di essere categorie separate.

C'è l'idea, c'è l'uso, c'è il codice, c'è la musica, c'è l'archivio, c'è il gesto ripetuto abbastanza volte da diventare mestiere e quello inatteso che apre un'altra strada.

**OraTorio... OraTorno...**

I puntini sono sospensione: non una fermata, ma quell'attimo in cui il pendolo attraversa il centro e puoi **guardare** prima che il moto prosegua.

Il pendolo continua.

> **Cartello d'uscita**  
> Se una cosa torna, non è detto che stia tornando indietro.  
> Potrebbe stare facendo un altro giro della spirale.

---

**ShiduLab 2002–2026**  
*Josta + G0gn4 + Scriba*
