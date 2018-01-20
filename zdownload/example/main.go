package main

import (
	"github.com/zouhuigang/package/zdownload"
	"os"
	"path"
)

func main() {
	//http://www.jctrans.com/tool/gjym.htm
	urlMap := [...]string{
		"ad.svg",
		"ae.svg",
		"af.svg",
		"ag.svg",
		"ai.svg",
		"al.svg",
		"am.svg",
		"ao.svg",
		"ar.svg",
		"at.svg",
		"au.svg",
		"az.svg",
		"bb.svg",
		"bd.svg",
		"be.svg",
		"bf.svg",
		"bg.svg",
		"bh.svg",
		"bi.svg",
		"bj.svg",
		"bl.svg",
		"bm.svg",
		"bn.svg",
		"bo.svg",
		"br.svg",
		"bs.svg",
		"bw.svg",
		"by.svg",
		"bz.svg",
		"ca.svg",
		"cf.svg",
		"cg.svg",
		"ch.svg",
		"ck.svg",
		"cl.svg",
		"cm.svg",
		"cn.svg",
		"co.svg",
		"cr.svg",
		"cs.svg",
		"cu.svg",
		"cy.svg",
		"cz.svg",
		"de.svg",
		"dj.svg",
		"dk.svg",
		"do.svg",
		"dz.svg",
		"ec.svg",
		"ee.svg",
		"eg.svg",
		"es.svg",
		"et.svg",
		"fi.svg",
		"fj.svg",
		"fr.svg",
		"ga.svg",
		"gb.svg",
		"gd.svg",
		"ge.svg",
		"gf.svg",
		"gh.svg",
		"gi.svg",
		"gm.svg",
		"gn.svg",
		"gr.svg",
		"gt.svg",
		"gu.svg",
		"gy.svg",
		"hk.svg",
		"hn.svg",
		"ht.svg",
		"hu.svg",
		"id.svg",
		"ie.svg",
		"il.svg",
		"in.svg",
		"iq.svg",
		"ir.svg",
		"is.svg",
		"it.svg",
		"jm.svg",
		"jo.svg",
		"jp.svg",
		"ke.svg",
		"kg.svg",
		"kh.svg",
		"kp.svg",
		"kr.svg",
		"kt.svg",
		"kw.svg",
		"kz.svg",
		"la.svg",
		"lb.svg",
		"lc.svg",
		"li.svg",
		"lk.svg",
		"lr.svg",
		"ls.svg",
		"lt.svg",
		"lu.svg",
		"lv.svg",
		"ly.svg",
		"ma.svg",
		"mc.svg",
		"md.svg",
		"mg.svg",
		"ml.svg",
		"mm.svg",
		"mn.svg",
		"mo.svg",
		"ms.svg",
		"mt.svg",
		"mu.svg",
		"mv.svg",
		"mw.svg",
		"mx.svg",
		"my.svg",
		"mz.svg",
		"na.svg",
		"ne.svg",
		"ng.svg",
		"ni.svg",
		"nl.svg",
		"no.svg",
		"np.svg",
		"nr.svg",
		"nz.svg",
		"om.svg",
		"pa.svg",
		"pe.svg",
		"pf.svg",
		"pg.svg",
		"ph.svg",
		"pk.svg",
		"pl.svg",
		"pr.svg",
		"pt.svg",
		"py.svg",
		"qa.svg",
		"ro.svg",
		"ru.svg",
		"sa.svg",
		"sb.svg",
		"sc.svg",
		"sd.svg",
		"se.svg",
		"sg.svg",
		"si.svg",
		"sk.svg",
		"sl.svg",
		"sm.svg",
		"sn.svg",
		"so.svg",
		"sr.svg",
		"st.svg",
		"sv.svg",
		"sy.svg",
		"sz.svg",
		"td.svg",
		"tg.svg",
		"th.svg",
		"tj.svg",
		"tm.svg",
		"tn.svg",
		"to.svg",
		"tr.svg",
		"tt.svg",
		"tw.svg",
		"tz.svg",
		"ua.svg",
		"ug.svg",
		"us.svg",
		"uy.svg",
		"uz.svg",
		"vc.svg",
		"ve.svg",
		"vn.svg",
		"ye.svg",
		"yu.svg",
		"za.svg",
		"zm.svg",
		"zr.svg",
		"zw.svg",
	}

	for _, v := range urlMap {

		res := zdownload.Img("https://www.qqzeng.com/ip/flags/4x3/" + v)
		//写入文件
		fileName := v
		os.MkdirAll(path.Dir(fileName), os.ModePerm)
		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		//写入图信息
		_, err = f.Write([]byte(res))
	}

}
