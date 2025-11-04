package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	logDir := "./logs"
	days := 7

	fmt.Printf("🔍 Recherche des logs de plus de %d jours dans %s...\n", days, logDir)

	// Créer le dossier de test
	os.MkdirAll(logDir, 0755)

	// Créer des fichiers de test
	createTestFiles(logDir)

	cutoffTime := time.Now().AddDate(0, 0, -days)
	deletedCount := 0

	entries, err := os.ReadDir(logDir)
	if err != nil {
		fmt.Printf("❌ Erreur: %v\n", err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".log" {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			if info.ModTime().Before(cutoffTime) {
				filepath := filepath.Join(logDir, entry.Name())
				err := os.Remove(filepath)
				if err == nil {
					fmt.Printf("🗑️  Supprimé: %s\n", entry.Name())
					deletedCount++
				}
			}
		}
	}

	fmt.Printf("✅ Nettoyage terminé! %d fichiers supprimés.\n", deletedCount)
}

func createTestFiles(logDir string) {
	// Implémentation similaire à Python...
}
