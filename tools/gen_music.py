import mido

mid = mido.MidiFile(ticks_per_beat=480)
track = mido.MidiTrack()
mid.tracks.append(track)
bpm = 140
track.append(mido.MetaMessage("set_tempo", tempo=mido.bpm2tempo(bpm)))
track.append(mido.Message("program_change", channel=0, program=80, time=0))  # lead square
track.append(mido.Message("program_change", channel=1, program=32, time=0))  # bass acoustic bass... use 0 piano-> use triangle via 80? use program 33
track.append(mido.Message("program_change", channel=2, program=80, time=0))

beat = 480
b16 = beat // 4

def note(msg_type, ch, pitch, vel, ticks):
    track.append(mido.Message(msg_type, channel=ch, note=pitch, velocity=vel, time=ticks))

chords = [57, 53, 55, 52]  # Am F G E
bass_root = [45, 41, 43, 40]

bars = 8
for bar in range(bars):
    chord = chords[bar % 4]
    root = bass_root[bar % 4]
    # bass: triangle, whole notes per bar (2 per bar)
    note("note_on", 1, root, 90, 0)
    note("note_off", 1, root, 0, beat * 2)
    note("note_on", 1, root, 90, 0)
    note("note_off", 1, root, 0, beat * 2)
    # arpeggio lead, square, 16ths
    arp = [chord, chord + 4, chord + 7, chord + 12]
    for i in range(16):
        p = arp[i % 4]
        note("note_on", 0, p + 12, 70 if i % 4 == 0 else 50, 0)
        note("note_off", 0, p + 12, 0, b16)

# melody on a second square channel, simple phrase
mel = [69, 72, 76, 72, 74, 72, 69, 67, 65, 69, 72, 69, 67, 64, 62, 64]
track2 = mido.MidiTrack()
mid.tracks.append(track2)
track2.append(mido.MetaMessage("set_tempo", tempo=mido.bpm2tempo(bpm)))
for bar in range(bars):
    for i in range(4):
        p = mel[(bar * 4 + i) % 16]
        track2.append(mido.Message("note_on", channel=2, note=p, velocity=85, time=0))
        track2.append(mido.Message("note_off", channel=2, note=p, velocity=0, time=beat))

mid.save("/tmp/music.mid")
print("midi saved")
