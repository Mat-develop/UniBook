import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
from matplotlib.patches import FancyBboxPatch

fig, ax = plt.subplots(figsize=(11, 15))
ax.set_xlim(0, 11)
ax.set_ylim(0, 15)
ax.axis('off')
fig.patch.set_facecolor('white')
ax.set_facecolor('white')

# ── helpers ──────────────────────────────────────────────────────────────────

def box(x, y, w, h, text, fc='#f8fafc', ec='#475569', fs=9.5, fw='normal',
        tc='#1e293b', ls='-', lw=1.8, pad=0.18):
    r = FancyBboxPatch((x, y), w, h, boxstyle=f"round,pad={pad}",
                       facecolor=fc, edgecolor=ec, linewidth=lw,
                       linestyle=ls, zorder=2)
    ax.add_patch(r)
    ax.text(x + w / 2, y + h / 2, text, ha='center', va='center',
            fontsize=fs, fontweight=fw, color=tc, zorder=3,
            multialignment='center', linespacing=1.5)

def arr(x1, y1, x2, y2, color='#475569', lw=1.8, style='->'):
    ax.annotate('', xy=(x2, y2), xytext=(x1, y1),
                arrowprops=dict(arrowstyle=style, color=color, lw=lw), zorder=4)

def label(x, y, text, fs=8, color='#64748b', ha='center', fw='normal', style='normal'):
    ax.text(x, y, text, ha=ha, va='center', fontsize=fs,
            color=color, fontweight=fw, style=style)

# ── title ─────────────────────────────────────────────────────────────────────

label(5.5, 14.5, 'Arquitetura Back-end — UniBook', fs=14, color='#0f172a', fw='bold')

# ── frontend ──────────────────────────────────────────────────────────────────

box(2.8, 13.1, 5.4, 0.95, 'Front-end\nReact / TypeScript',
    fc='#dbeafe', ec='#2563eb', fs=10.5, fw='bold', tc='#1d4ed8', lw=2.5)

arr(5.5, 13.1, 5.5, 12.5, color='#334155')
label(6.9, 12.8, 'HTTP Request + Bearer Token', fs=8, color='#64748b')

# ── server container ──────────────────────────────────────────────────────────

srv = FancyBboxPatch((0.4, 2.3), 10.2, 10.0,
                     boxstyle="round,pad=0.25",
                     facecolor='#f1f5f9', edgecolor='#94a3b8', linewidth=2,
                     linestyle='--', zorder=1)
ax.add_patch(srv)
label(5.5, 12.25, 'Servidor Go  ·  Gorilla Mux', fs=10, color='#475569', fw='bold', style='normal')

# ── middleware row ────────────────────────────────────────────────────────────

MW_Y = 10.9
box(0.7,  MW_Y, 2.1, 0.9, 'CORS\nMiddleware',
    fc='#fef9c3', ec='#ca8a04', fs=9, fw='bold', tc='#78350f')
box(3.1,  MW_Y, 2.1, 0.9, 'Logger\nMiddleware',
    fc='#fef9c3', ec='#ca8a04', fs=9, fw='bold', tc='#78350f')
box(5.5,  MW_Y, 3.2, 0.9, 'IsAuth  ·  JWT HS256',
    fc='#fee2e2', ec='#dc2626', fs=9, fw='bold', tc='#991b1b')

arr(2.8,  MW_Y + 0.45, 3.1,  MW_Y + 0.45)
arr(5.2,  MW_Y + 0.45, 5.5,  MW_Y + 0.45)

# 401 side note
arr(8.7, MW_Y + 0.45, 9.8, MW_Y + 0.45, color='#dc2626')
label(10.15, MW_Y + 0.45, '401', fs=9, color='#dc2626', fw='bold')

arr(5.5, MW_Y, 5.5, 10.1, color='#334155')   # down to routes

# ── routes ────────────────────────────────────────────────────────────────────

box(0.7, 9.25, 8.5, 0.8,
    '/login     /users     /post     /c     /tag',
    fc='#ede9fe', ec='#7c3aed', fs=9.5, tc='#4c1d95')
label(0.82, 10.1, 'Rotas', fs=8, color='#7c3aed', ha='left', style='italic')

arr(5.5, 9.25, 5.5, 8.65, color='#334155')

# ── layers ────────────────────────────────────────────────────────────────────

LAYER_X = 0.7
LAYER_W = 8.5
LAYER_H = 1.05
GAP     = 0.38

# Handler
H_Y = 7.65
box(LAYER_X, H_Y, LAYER_W, LAYER_H,
    'Handler\nRecebe a requisição · extrai userId do JWT · retorna JSON',
    fc='#dcfce7', ec='#16a34a', fs=9.5, fw='bold', tc='#14532d')

label(9.6, H_Y + LAYER_H/2, 'User\nPost\nCommunity\nTag',
      fs=7.5, color='#94a3b8', ha='left')

arr(5.5, H_Y, 5.5, H_Y - GAP, color='#334155')

# Service
S_Y = H_Y - GAP - LAYER_H
box(LAYER_X, S_Y, LAYER_W, LAYER_H,
    'Service\nRegras de negócio · validações de domínio',
    fc='#bbf7d0', ec='#15803d', fs=9.5, fw='bold', tc='#14532d')

arr(5.5, S_Y, 5.5, S_Y - GAP, color='#334155')

# Repository
R_Y = S_Y - GAP - LAYER_H
box(LAYER_X, R_Y, LAYER_W, LAYER_H,
    'Repository\nPrepared statements SQL · mapeamento para structs Go',
    fc='#86efac', ec='#166534', fs=9.5, fw='bold', tc='#14532d')

DB_GAP = 0.75
arr(5.5, R_Y, 5.5, R_Y - DB_GAP + 0.1, color='#334155')
label(7.0, R_Y - DB_GAP / 2, 'Queries SQL\nResultados', fs=8, color='#64748b')

# ── database ──────────────────────────────────────────────────────────────────

DB_Y = R_Y - DB_GAP - 1.1
box(2.5, DB_Y, 6.0, 1.1,
    'MySQL\nusuários · posts · comunidades · tags',
    fc='#ffedd5', ec='#c2410c', fs=10, fw='bold', tc='#7c2d12', lw=2.5)

# ── save ──────────────────────────────────────────────────────────────────────

plt.savefig('arquitetura_backend.png', dpi=160, bbox_inches='tight',
            facecolor='white', edgecolor='none')
print('Imagem salva: arquitetura_backend.png')
