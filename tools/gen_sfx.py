import struct, math, random

random.seed(7)
SR = 22050

def env(i, n, a=0.005, r=0.3):
    t = i / n
    if t < a: return t / a
    return max(0.0, 1.0 - (t - a) / (1 - a)) ** 1.5

def tone(freq, dur, kind="sine", vol=0.5, slide_to=None):
    n = int(SR * dur)
    out = []
    for i in range(n):
        t = i / SR
        f = freq if slide_to is None else freq + (slide_to - freq) * (i / n)
        ph = 2 * math.pi * f * t
        if kind == "sine": v = math.sin(ph)
        elif kind == "square": v = 1.0 if math.sin(ph) >= 0 else -1.0
        elif kind == "saw": v = 2 * (f * t % 1.0) - 1
        else: v = math.sin(ph)
        out.append(v * vol * env(i, n))
    return out

def mix(*tracks):
    n = max(len(t) for t in tracks)
    return [sum(t[i] if i < len(t) else 0 for t in tracks) for i in range(n)]

def shift(track, delay):
    return [0.0] * int(SR * delay) + track

def save_wav(path, data):
    n = len(data)
    with open(path, "wb") as f:
        f.write(b"RIFF" + struct.pack("<I", 36 + n * 2) + b"WAVEfmt " + struct.pack("<IHHIIHH", 16, 1, 1, SR, SR * 2, 2, 16) + b"data" + struct.pack("<I", n * 2))
        for v in data:
            f.write(struct.pack("<h", max(-32767, min(32767, int(v * 32767)))))

# move: short soft tick
save_wav("/tmp/sfx/move.wav", mix(tone(196, 0.06, "square", 0.25, slide_to=150)))
# eat: bright two-note ding
save_wav("/tmp/sfx/eat.wav", mix(tone(880, 0.12, "sine", 0.5), shift(tone(1318, 0.2, "sine", 0.5), 0.08)))
# over: descending saw
save_wav("/tmp/sfx/over.wav", tone(380, 0.7, "saw", 0.4, slide_to=55))
# start: quick ascending arp
save_wav("/tmp/sfx/start.wav", mix(tone(440, 0.1, "square", 0.3), shift(tone(660, 0.1, "square", 0.3), 0.09), shift(tone(880, 0.22, "square", 0.3), 0.18)))
print("sfx done")
