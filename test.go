

    r.GET("/new-wallet", func(c *gin.Context) {
        wallet := NewWallet()
		a := wallet.GetAddress()
		wallets[a] = wallet
        c.JSON(http.StatusOK, gin.H{"wallets": wallets})
    })

