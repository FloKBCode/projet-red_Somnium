# 📘 Document de Gestion de Projet - Somnium 🌙

## 1. Présentation du Projet

### 1.1 Concept
**Somnium** est un RPG textuel développé en Go dans le cadre du Projet RED.  
Le jeu plonge le joueur dans un univers onirique où il incarne un aventurier explorant les méandres du rêve et de la réalité.

### 1.2 Objectifs Pédagogiques
- Maîtriser les structures de données et méthodes en Go  
- Implémenter un système de jeu complet (personnage, combat, économie)  
- Collaborer efficacement en équipe de développement  
- Appliquer les bonnes pratiques de développement logiciel  

---

## 2. Organisation de l'Équipe

### 2.1 Répartition des Rôles

| Développeur | Rôle Principal              | Fichiers Assignés                     |
|-------------|-----------------------------|---------------------------------------|
| **Florence** (Dev 1) | Architecture Core & Économie | `main.go`, `character/`, `shop/forge.go` |
| **Sarah** (Dev 2)    | Systèmes de Combat          | `combat/`, `character/inventory.go`   |
| **Marly** (Dev 3)    | Interface & UX              | `shop/merchant.go`, utilitaires       |
| **Ananya**           | Contribution spécifique     | Système d'Initiative                  |

### 2.2 Méthodes de Collaboration
- **Planning détaillé** sur 4 jours avec répartition horaire  
- **Attribution claire** des fichiers par développeur  
- **Intégration progressive** des modules développés  

---

## 3. Compréhension des Consignes

### 3.1 Fonctionnalités demandées

#### 🧍 Système de Personnage
- ✅ Création de personnage avec validation du nom (uniquement des lettres)  
- ✅ 3 races de base : Humain, Elfe, Nain (+2 ajoutées : Spectre, Abysse)  
- ✅ Stats de base : HP, Mana, Inventaire, Argent  
- ✅ Système d'équipement avec 3 emplacements (Tête, Torse, Pieds)  

#### 🎒 Système d'Inventaire
- ✅ Limite de 10 objets au départ (améliorable avec l’argent)  
- ✅ Potions de vie (+50 PV) et poison (-10 PV/sec pendant 3 sec)  
- ✅ Possibilité d’agrandir l’inventaire  

#### 💰 Système Économique
- ✅ Marchand avec au moins 8 objets  
- ✅ Forgeron avec au moins 3 recettes d’artisanat  
- ✅ Gestion de l’argent et des matériaux pour le craft  

#### ⚔️ Système de Combat
- ✅ Combat d’entraînement contre un Gobelin  
- ✅ 2 sorts minimum : Coup de poing (8 dégâts, 5 mana), Boule de feu (18 dégâts, 15 mana)  
- ✅ Pattern d’attaque du Gobelin (attaque normale + attaque spéciale tous les 3 tours)  
- ✅ Système d’initiative (implémenté par Ananya)  

### 3.2 Missions Bonus réalisées
- ✅ Mission 1 – Initiative (ordre du combat)  
- ✅ Mission 2 – Système d’XP et montée de niveau  
- ✅ Mission 3 – Combat magique étendu (Soin, Bouclier, Chaîne d’éclairs, etc.)  
- ✅ Mission 4 – Gestion avancée de la mana (avec potions)  
- ✅ Mission 5 – Exploration de donjon (couches, événements aléatoires)  
- ✅ Mission 6 – Easter eggs (ABBA & Spielberg cachés dans le menu "Qui sont-ils")  

➡️ **Toutes les missions bonus ont été validées ! 🎉**

---

## 4. Planification et Suivi

### 4.1 Planning initial (4 jours)

#### 📅 Jour 1 – Fondations
- Matin : Architecture de base, structures `Character`  
- Après-midi : Création personnage, menu principal, combat de base  

#### 📅 Jour 2 – Systèmes avancés
- Matin : Économie (forge, marchand)  
- Après-midi : Combat avancé, système d’équipement  

#### 📅 Jour 3 – Missions Bonus
- Matin : XP, système de mana  
- Après-midi : Donjons, sorts avancés  

#### 📅 Jour 4 – Finalisation
- Matin : Code review, optimisation  
- Après-midi : Préparation à l’oral, tests finaux  

### 4.2 État d’avancement

#### ✅ Terminé
- Structure de base du projet (dossiers, go.mod)  
- Création de personnages (5 races, 4 classes)  
- Menu principal fonctionnel  
- Système d’inventaire avec extension  
- Interface du marchand et forge (recettes)  
- Combat contre gobelins et autres monstres  
- Initiative en combat  
- Système d’XP et montée de niveau  
- Exploration des couches du Labyrinthe  
- Sauvegarde/chargement de la partie  
- Système de quêtes  

#### ⚠️ Presque fini
- Forge (recettes OK mais équilibrage en cours)  
- Équipements (bonus parfois mal appliqués)  
- Sorts avancés (certains bugs mineurs)  
- Interface utilisateur (affichage CLI à améliorer)  

#### 🔄 En cours
- Correction des derniers bugs combat/inventaire  
- Amélioration visuelle des menus et messages  

---

## 5. Problèmes rencontrés

### 5.1 Techniques

#### 🎒 Inventaire
- ❌ Bug : l’inventaire affichait "plein" même quand il restait de la place  
- ✅ Solution : refactor des fonctions `HasInventorySpace()`, `TakeItem()` et `UseItem()`  

#### ⚔️ Combat
- ❌ Bug : affichage chaotique, plantages fréquents  
- ✅ Solution : création d’une struct `CombatState` + nettoyage de l’écran à chaque tour  

#### 🟨 Go (langage)
- ❌ Erreurs fréquentes avec la syntaxe des structs/méthodes  
- ✅ Solution : adoption stricte du camelCase et apprentissage via tutos  

### 5.2 Organisation

#### 📂 Structure des fichiers
- ❌ Mauvaise gestion des imports et packages  
- ✅ Solution : refonte des dossiers + configuration correcte de `go.mod`  

#### 📝 Documentation
- ❌ README et doc projet oubliés au début  
- ✅ Solution : création de ce document + ajout de commentaires dans le code  

---

## 6. Bilan pédagogique

### 6.1 Ce qui a bien marché
- Respect du planning initial  
- Répartition claire des rôles  
- Réalisation de toutes les missions bonus  
- Bonne entraide entre membres  

### 6.2 Ce qu’on referait autrement
- Approfondir Go avant le projet  
- Tester plus tôt les fonctionnalités  
- Documenter au fur et à mesure  
- Mieux estimer le temps pour certaines features  

### 6.3 Apprentissages clés
1. Faire des prototypes rapides avant de coder entièrement  
2. Importance des tests unitaires  
3. Communication claire sur les interfaces entre modules  
4. Capacité à s’adapter et revoir le plan  

---

## 7. Prochaines étapes

### 7.1 Corrections
- [ ] Finaliser l’équilibrage forge et équipements  
- [ ] Corriger les derniers bugs d’inventaire et de combat  
- [ ] Nettoyer l’affichage CLI  

### 7.2 Préparation orale
- [ ] Démo jouable (5-10 minutes)  
- [ ] Réponses aux questions techniques (Go, architecture, choix)  
- [ ] Justification du thème choisi  

### 7.3 Livraison finale
- [ ] Code propre et rangé  
- [ ] README complet (installation + gameplay)  
- [ ] Ce document de gestion de projet  
- [ ] Préparation de la présentation finale  

---

## 🏁 Conclusion

Le projet **Somnium** a été un vrai défi : découverte du langage Go, gestion en équipe de 4, et mise en place d’un RPG complet en CLI.  
Nous avons réussi à **couvrir toutes les fonctionnalités demandées et bonus**, malgré quelques bugs mineurs restants.  

Ce projet nous a appris l’importance de :  
- Planifier et s’organiser,  
- Communiquer efficacement,  
- Tester régulièrement,  
- Documenter son travail.  

Nous sommes fiers d’avoir réalisé un jeu jouable et immersif, et espérons que le jury appréciera notre approche créative.  

---

*Document rédigé par l’équipe Somnium*  
👩‍💻 Florence – Sarah – Marly – Ananya  
