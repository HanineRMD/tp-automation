#!/bin/bash

# Exercice 1: Nettoyer les fichiers logs de plus de 7 jours

LOG_DIR="./logs"
DAYS=7

echo "🔍 Recherche des logs de plus de $DAYS jours dans $LOG_DIR..."

# Créer le dossier de test
mkdir -p "$LOG_DIR"

# Créer des fichiers de test
create_test_files() {
    for i in {0..4}; do
        if [ $i -lt 2 ]; then
            age_days=$((RANDOM % 4))  # 0-3 jours
        else
            age_days=$((8 + RANDOM % 8))  # 8-15 jours
        fi
        
        filename="test_log_$i.log"
        filepath="$LOG_DIR/$filename"
        
        echo "Fichier log de test $i" > "$filepath"
        
        # Modifier la date du fichier
        file_time=$(date -d "-$age_days days" +%s)
        touch -d "@$file_time" "$filepath"
        
        date_str=$(date -d "@$file_time" "+%Y-%m-%d")
        echo "📁 Créé: $filename (âge: $age_days jours, date: $date_str)"
    done
}

create_test_files

echo ""
echo "🗑️  Suppression des fichiers de plus de $DAYS jours..."
find "$LOG_DIR" -name "*.log" -type f -mtime "+$DAYS" -delete -print

echo "✅ Nettoyage terminé!"