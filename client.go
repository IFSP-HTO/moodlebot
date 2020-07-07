package main

import (
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func main() {
	// Set up authentication information.
	auth := sasl.NewPlainClient("", "", "")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"admin@localhost"}

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

	raw = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=ifsp.edu.br;
	s=F67E9BF8-6F46-11E4-B529-30183D1F85CC; t=1594086664;
	bh=lLHSzWYuSPUpHndn4qH8898q2H7Id0yAv/Or7iqytd0=;
	h=Date:To:From:Reply-To:Subject:Message-ID:MIME-Version:
	 Content-Type;
	b=ypf6JklC4Uq7/NVXsA2t2K+4S/svO/sVg74CGriejXbGgQlC03UFZxWrkjNI+BQz1
	 zoGAvGdLosIK0Mh0OLE7AISx585oFrgYK+dUoqLbhs+pwLg5lB/YkIL8kNlrvnGE03
	 6+nwVp1beltwpFVdWQq5RUJfNgLrmtw1eSgCl3Fw=
X-Virus-Scanned: amavisd-new at ifsp.edu.br
Received: from email.ifsp.edu.br ([127.0.0.1])
	by localhost (email.ifsp.edu.br [127.0.0.1]) (amavisd-new, port 10026)
	with ESMTP id avrO7PBmEgtq for <prof.flaviobarros@gmail.com>;
	Mon,  6 Jul 2020 22:51:04 -0300 (BRT)
Received: from zorro18 (unknown [200.133.218.108])
	by email.ifsp.edu.br (Postfix) with ESMTPSA id 7DD1317B0C3
	for <prof.flaviobarros@gmail.com>; Mon,  6 Jul 2020 22:51:04 -0300 (BRT)
Date: Mon, 6 Jul 2020 22:51:03 -0300
To: "prof.flaviobarros@gmail.com" <prof.flaviobarros@gmail.com>
From: "Administrador CTI (via Moodle IFSP HTO)" <naoresponder.hto@ifsp.edu.br>
Reply-To: "=?utf-8?B?TsOjbyByZXNwb25kYSBhIGVzdGEgbWVuc2FnZW0u?=" <naoresponder.hto@ifsp.edu.br>
Subject: =?utf-8?B?VXN1w6FyaW8gYWNlc3NvdSBjdXJzbyBNb29kbGU=?=
Message-ID: <5f03d507545475.22811242/ead@hto.ifsp.edu.br>
X-Mailer: PHPMailer 6.1.3 (https://github.com/PHPMailer/PHPMailer)
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8


2061427 - Kenia Cristina Pereira Silva entrou no curso Moodle
para Professores.`

	raw = `MIME-Version: 1.0
	Date: Tue, 7 Jul 2020 17:37:35 -0300
	Sender: prof.flaviobarros@gmail.com
	X-Google-Sender-Delegation: prof.flaviobarros@gmail.com
	Message-ID: <CAHJVo3XLBW3CEViMEKriS6YvzRyLuiQo5Hck3At_Le_9ff=cdA@mail.gmail.com>
	Subject: Teste
	From: Flavio Barros <prof.flaviobarros@gmail.com>
	To: admin@stat-math.com.br
	Content-Type: multipart/alternative; boundary="00000000000027cde805a9dff699"
	
	--00000000000027cde805a9dff699
	Content-Type: text/plain; charset="UTF-8"
	
	Teste do e-mail prof.flaviobarros@gmail.com
	
	--00000000000027cde805a9dff699
	Content-Type: text/html; charset="UTF-8"
	
	<div dir="ltr">Teste do e-mail <a href="mailto:prof.flaviobarros@gmail.com">prof.flaviobarros@gmail.com</a><br></div>
	
	--00000000000027cde805a9dff699--`

	msg := strings.NewReader(raw)
	err := smtp.SendMail("localhost:2525", auth, "prof.flaviobarros@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
