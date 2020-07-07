package main

import (
	"fmt"
	"strings"

	"github.com/jhillyerd/enmime"
)

func main() {

	raw := `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
	d=gmail.com; s=20161025;
	h=mime-version:sender:from:date:message-id:subject:to;
	bh=r6aEXPJxlkJvhIzfkF8UR9cekIJD5UyldQM9fu8tcSE=;
	b=p5kGoUPtXuRavHzsObu7NGJvc9pCrIJkVmQ9hftUHFk394dwkrfFcwkNk4Y2/CWvGb
	 RonAJsThIujOV78M5SBIpBObxleKvGSWr3Kb+upjHzqBDUQNxBDaItbQvP8+Ivarywiu
	 a5YjyOPyk69Z2VUbqJd5hVFRGZMdU9GWpmmZ7Z3r28MHvRmvJObEcG+6qVDBPImT4Y+I
	 dKJbX0RIHln4DYsxBuSg1cBx1n3GvocoI+0lSqLqhCFJ2ysVUesCEPZtskikoFs7w3O2
	 +AWUnl09OioVY3JOZ5RoJF81bD0eYXN3xKVc2ZwVgZk0XxD35twhnX4vws6DSty1Wx++
	 ofLg==
X-Google-DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
	d=1e100.net; s=20161025;
	h=x-gm-message-state:mime-version:sender:from:date:message-id:subject
	 :to;
	bh=r6aEXPJxlkJvhIzfkF8UR9cekIJD5UyldQM9fu8tcSE=;
	b=lqwH1ueldIKbz2rn0iBFPdd7SocaOTzXGDjyMcAj+bbNMXwBHuYhiJtrP8aucNmPFT
	 +VmRD3L5ZMwWYbLJtttvmdvtjJUnRXVXc8yPM7QPipcRsy1GtAIGu7EzUzzlY5WHnXft
	 YZQ5PNt2a1pwhXMkxrKNTmQ16xtP333+yyve709g7rEl75J336nLvL14b8OG9HcUXnum
	 34BtYKpDM3e7tecptkVC8a8jq/Wd5K+EzyXzZWgX6cg3i+8bIQFs0t/2MkUTCR/e5q/C
	 4j1B0zZkn2Jn1XRilxCwpLFiAeDnEqPs43RFujwCR4luqBWXeEcpq4QIPwPiD7ONRm6V
	 1bZQ==
X-Gm-Message-State: AOAM532qynVsDzhTY4O1nAZ5pdYdvQianjWu6/dc75qhV7p0qAirUcwQ
7zu55zmrfWhTxOWerDWKf7u+9f7DYp2o0d+Wed9Fyv66
X-Google-Smtp-Source: ABdhPJziFIjFCITRLJsMrelkIePZ2j2Y49Tq9aGeTo4K//xZrRZwJPHq5IRzGykstvc3ReBmxGgyqBnectds+5BOHU0=
X-Received: by 2002:a05:6102:812:: with SMTP id g18mr15709058vsb.131.1592125279176;
Sun, 14 Jun 2020 02:01:19 -0700 (PDT)
MIME-Version: 1.0
Sender: flaviomargarito@gmail.com
X-Google-Sender-Delegation: flaviomargarito@gmail.com
From: Flavio Barros <prof.flaviobarros@gmail.com>
Date: Sun, 14 Jun 2020 06:01:07 -0300
X-Google-Sender-Auth: O59WCFnZ1rrn8ni-kRrZcE62XzU
Message-ID: <CAHJVo3Uw4g2FY_pe10VtEaeErT=X8zCdjMBpEyTN_Lx7YT602w@mail.gmail.com>
Subject: Email
To: admin@stat-math.com.br
Content-Type: multipart/alternative; boundary="000000000000b40a4d05a8078d01"

--000000000000b40a4d05a8078d01
Content-Type: text/plain; charset="UTF-8"

Somente um teste.

--000000000000b40a4d05a8078d01
Content-Type: text/html; charset="UTF-8"

<div dir="ltr">Somente um teste.<br></div>

--000000000000b40a4d05a8078d01--
`

	r := strings.NewReader(raw)
	env, err := enmime.ReadEnvelope(r)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(env.Text)
	for k, v := range env.Root.Header {
		fmt.Printf("Key: %s; Value: %s\n", k, v)
	}

	fmt.Printf("%T", env.Root.Header.Get("Sender"))

}
