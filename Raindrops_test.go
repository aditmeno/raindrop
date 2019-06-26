package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Raindrops", func() {
	Describe("Print Pling Plang Plong based on whether the number passed in the URL is divisible by 3, 5 or 7", func() {

		Context("Print Pling if number is divisible by 3", func() {
			It("Should return Pling", func() {
				req, _ := http.NewRequest("GET", "/3", nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(RainDrops)
				handler.ServeHTTP(rr, req)
				Expect(rr.Body.String()).To(Equal("Pling"))
			})
		})

		Context("Print Plang if number is divisible by 5", func() {
			It("Should return Plang", func() {
				req, _ := http.NewRequest("GET", "/5", nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(RainDrops)
				handler.ServeHTTP(rr, req)
				Expect(rr.Body.String()).To(Equal("Plang"))
			})
		})

		Context("Print Plong if number is divisible by 7", func() {
			It("Should return Plong", func() {
				req, _ := http.NewRequest("GET", "/7", nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(RainDrops)
				handler.ServeHTTP(rr, req)
				Expect(rr.Body.String()).To(Equal("Plong"))
			})
		})

		Context("Print PlingPlangPlong if number is divisible by 3, 5 and 7", func() {
			It("Should return PlingPlangPlong", func() {
				req, _ := http.NewRequest("GET", "/105", nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(RainDrops)
				handler.ServeHTTP(rr, req)
				Expect(rr.Body.String()).To(Equal("PlingPlangPlong"))
			})
		})

		Context("Print PlingPlang if number is divisible by 3 and 5", func() {
			It("Should return PlingPlang", func() {
				req, _ := http.NewRequest("GET", "/15", nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(RainDrops)
				handler.ServeHTTP(rr, req)
				Expect(rr.Body.String()).To(Equal("PlingPlang"))
			})
		})

		Context("Print PlingPlong if number is divisible by 3 and 7", func() {
			It("Should return PlingPlong", func() {
				req, _ := http.NewRequest("GET", "/21", nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(RainDrops)
				handler.ServeHTTP(rr, req)
				Expect(rr.Body.String()).To(Equal("PlingPlong"))
			})
		})

		Context("Print PlangPlong if number is divisible by 5 and 7", func() {
			It("Should return PlangPlong", func() {
				req, _ := http.NewRequest("GET", "/35", nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(RainDrops)
				handler.ServeHTTP(rr, req)
				Expect(rr.Body.String()).To(Equal("PlangPlong"))
			})
		})
	})
})
