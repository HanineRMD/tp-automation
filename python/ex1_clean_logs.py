#!/usr/bin/env python3
"""
Exercice 1: Nettoyer les fichiers logs de plus de 7 jours
"""

import os
import time
import argparse

def clean_old_logs(log_dir="./logs", days=7):
    """Supprime les fichiers logs plus anciens que X jours"""
    print(f"🔍 Recherche des logs de plus de {days} jours dans {log_dir}...")
    
    # Créer le dossier de test s'il n'existe pas
    os.makedirs(log_dir, exist_ok=True)
    
    # Créer des fichiers de test
    create_test_files(log_dir)
    
    cutoff_time = time.time() - (days * 24 * 60 * 60)
    deleted_count = 0
    
    try:
        for filename in os.listdir(log_dir):
            filepath = os.path.join(log_dir, filename)
            if os.path.isfile(filepath) and filename.endswith('.log'):
                file_time = os.path.getmtime(filepath)
                if file_time < cutoff_time:
                    os.remove(filepath)
                    print(f"🗑️  Supprimé: {filename}")
                    deleted_count += 1
        
        print(f"✅ Nettoyage terminé! {deleted_count} fichiers supprimés.")
        
    except Exception as e:
        print(f"❌ Erreur: {e}")

def create_test_files(log_dir):
    """Crée des fichiers de test pour la démo"""
    import random
    from datetime import datetime, timedelta
    
    # Créer 5 fichiers avec différentes dates
    for i in range(5):
        filename = f"test_log_{i}.log"
        filepath = os.path.join(log_dir, filename)
        
        # Certains fichiers récents, d'autres anciens
        if i < 2:  # Fichiers récents
            age_days = random.randint(0, 3)
        else:      # Fichiers anciens
            age_days = random.randint(8, 15)
            
        # Créer le fichier avec une date modifiée
        with open(filepath, 'w') as f:
            f.write(f"Fichier log de test {i}\n")
        
        # Modifier la date du fichier
        file_time = time.time() - (age_days * 24 * 60 * 60)
        os.utime(filepath, (file_time, file_time))
        
        date_str = datetime.fromtimestamp(file_time).strftime("%Y-%m-%d")
        print(f"📁 Créé: {filename} (âge: {age_days} jours, date: {date_str})")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Nettoyer les anciens fichiers logs')
    parser.add_argument('--dir', default='./logs', help='Dossier des logs')
    parser.add_argument('--days', type=int, default=7, help='Nombre de jours à conserver')
    
    args = parser.parse_args()
    clean_old_logs(args.dir, args.days)