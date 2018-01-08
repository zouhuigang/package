/*
分页库，配合前端和数据库使用,根据之前使用的PHP分页函数转化而来
*/
package zcore

import (
	"fmt"
	"math"
)

/*   分页函数
 *   currentPage    当前页
 *   totalRecords   总页数
 *   url            跳转链接
 *   pageSize       一页显示多少条数据
 */
func MultLink(currentPage int64, totalRecords int64, url string, pageSize int64) string {
	var counter int64
	if pageSize == 0 {
		pageSize = 10
	}
	//lang_prev := `&nbsp;`
	//lang_next := `&nbsp;`
	lang_prev := `<`
	lang_next := `>`

	if totalRecords <= pageSize { //总条数少于一页显示的数目,无需分页
		return ""
	}

	var mult string
	totalPages := math.Ceil(float64(totalRecords) / float64(pageSize))

	mult += `<div class="pages">`
	if currentPage < 1 {
		currentPage = 1
	}

	if currentPage > 1 {
		cp := fmt.Sprintf("%d", (currentPage - 1))
		mult += `<a href="` + url + `p=` + cp + `" class="page_icon">` + lang_prev + `</a>`
	} else {
		mult += `<b class="page_icon">` + lang_prev + `</b>`
	}

	if totalPages < 13 { //总页数少于13页
		for counter = 1; counter <= int64(totalPages); counter++ {
			counters := fmt.Sprintf("%d", counter)
			if counter == currentPage {
				mult += `<b>` + counters + `</b>`
			} else {
				mult += `<a href="` + url + `p=` + counters + `">` + counters + `</a>`
			}
		}

	} else if totalPages > 11 { //总页数大于11页
		if currentPage < 7 {
			for counter = 1; counter < 10; counter++ {
				counters := fmt.Sprintf("%d", counter)
				if counter == currentPage {
					mult += `<b>` + counters + `</b>`
				} else {
					mult += `<a href="` + url + `p=` + counters + `">` + counters + `</a>`
				}
			}
			totalPages1 := fmt.Sprintf("%d", (totalPages - 1))
			totalPages2 := fmt.Sprintf("%d", totalPages)
			mult += `<span>&#8230;</span><a href="` + url + `p=` + totalPages1 + `">` + totalPages1 + `</a>`
			mult += `<a href="` + url + `p=` + totalPages2 + `">` + totalPages2 + `</a>`
		} else if int64(totalPages-6) > currentPage && (currentPage > 6) {
			mult += `<a href="` + url + `p=1">1</a><a href="` + url + `p=2">2</a><span>&#8230;</span>`
			for counter = currentPage - 3; counter <= currentPage+3; counter++ {
				counters := fmt.Sprintf("%d", counter)
				if counter == currentPage {
					mult += `<b>` + counters + `</b>`
				} else {
					mult += `<a href="` + url + `p=` + counters + `">` + counters + `</a>`
				}
			}
			totalPages1 := fmt.Sprintf("%d", (totalPages - 1))
			totalPages2 := fmt.Sprintf("%d", totalPages)
			mult += `<span>&#8230;</span><a href="` + url + `p=` + totalPages1 + `">` + totalPages1 + `</a>`
			mult += `<a href="` + url + `p=` + totalPages2 + `">` + totalPages2 + `</a>`
		} else {
			mult += `<a href="` + url + `p=1">1</a><a href="` + url + `p=2">2</a><span>&#8230;</span>`
			for counter = int64(totalPages - 8); counter <= int64(totalPages); counter++ {
				counters := fmt.Sprintf("%d", counter)
				if counter == currentPage {
					mult += `<b>` + counters + `</b>`
				} else {
					mult += `<a href="` + url + `p=` + counters + `">` + counters + `</a>`
				}
			}
		}
	}

	if currentPage < (counter - 1) {
		currentPages := fmt.Sprintf("%d", (currentPage + 1))
		mult += `<a href="` + url + `p=` + currentPages + `" class="nextprev">` + lang_next + `</a>`
	} else {
		mult += `<b class="nextprev">` + lang_next + `</b>`
	}
	//$mult .= '<div class="fl">记录<strong style="color:red;">'.$totalRecords.'</strong>条&nbsp;&nbsp;共<strong style="color:red;">'.$totalPages.'</strong>页</div>';
	mult += `</div>`
	return mult

	/*
	   $lang_prev = '&nbsp;';
	   $lang_next = '&nbsp;';
	   if ($totalRecords <= $pageSize)
	       return '';
	   $mult = '';
	   $totalPages = ceil($totalRecords / $pageSize);
	   $mult .= '<div class="pages">';
	   $currentPage < 1 && $currentPage = 1;
	   if ($currentPage > 1) {
	       $mult .= '<a href="' . $url . 'p=' . ($currentPage - 1) . '" class="page_icon">' . $lang_prev . '</a>';
	   } else {
	       $mult .= '<b class="page_icon">' . $lang_prev . '</b>';
	   }
	   if ($totalPages < 13) {
	       for ($counter = 1; $counter <= $totalPages; $counter++) {
	           if ($counter == $currentPage) {
	               $mult .= '<b>' . $counter . '</b>';
	           } else {
	               $mult .= '<a href="' . $url . 'p=' . $counter . '">' . $counter . '</a>';
	           }
	       }
	   } elseif ($totalPages > 11) {
	       if ($currentPage < 7) {
	           for ($counter = 1; $counter < 10; $counter++) {
	               if ($counter == $currentPage) {
	                   $mult .= '<b>' . $counter . '</b>';
	               } else {
	                   $mult .= '<a href="' . $url . 'p=' . $counter . '">' . $counter . '</a>';
	               }
	           }
	           $mult .= '<span>&#8230;</span><a href="' . $url . 'p=' . ($totalPages - 1) . '">' . ($totalPages - 1) . '</a><a href="' . $url . 'p=' . $totalPages . '">' . $totalPages . '</a>';
	       } elseif ($totalPages - 6 > $currentPage && $currentPage > 6) {
	           $mult .= '<a href="' . $url . 'p=1">1</a><a href="' . $url . 'p=2">2</a><span>&#8230;</span>';
	           for ($counter = $currentPage - 3; $counter <= $currentPage + 3; $counter++) {
	               if ($counter == $currentPage) {
	                   $mult .= '<b>' . $counter . '</b>';
	               } else {
	                   $mult .= '<a href="' . $url . 'p=' . $counter . '">' . $counter . '</a>';
	               }
	           }
	           $mult .= '<span>&#8230;</span><a href="' . $url . 'p=' . ($totalPages - 1) . '">' . ($totalPages - 1) . '</a><a href="' . $url . 'p=' . $totalPages . '">' . $totalPages . '</a>';
	       } else {
	           $mult .= '<a href="' . $url . 'p=1">1</a><a href="' . $url . 'p=2">2</a><span>&#8230;</span>';
	           for ($counter = $totalPages - 8; $counter <= $totalPages; $counter++) {
	               if ($counter == $currentPage) {
	                   $mult .= '<b>' . $counter . '</b>';
	               } else {
	                   $mult .= '<a href="' . $url . 'p=' . $counter . '">' . $counter . '</a>';
	               }
	           }
	       }
	   }
	   if ($currentPage < $counter - 1) {
	       $mult .= '<a href="' . $url . 'p=' . ($currentPage + 1) . '" class="nextprev">' . $lang_next . '</a>';
	   } else {
	       $mult .= '<b class="nextprev">' . $lang_next . '</b>';
	   }
	   //$mult .= '<div class="fl">记录<strong style="color:red;">'.$totalRecords.'</strong>条&nbsp;&nbsp;共<strong style="color:red;">'.$totalPages.'</strong>页</div>';
	   $mult .= '</div>';
	   return $mult;*/
}
